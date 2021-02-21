package service

import (
	"context"
	"strings"
	"time"

	"github.com/luobin998877/go_grpc_with_consul/pb"
	"github.com/luobin998877/go_utility/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var sysError = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 44, 32, 34, 109, 115, 103, 34, 58, 34, 115, 121, 115, 32, 101, 114, 114, 111, 114, 34, 125}

// Client struct
type Client struct {
	c    pb.CommonClient
	conn *grpc.ClientConn
}

// CreateConnection with gRPC
func CreateConnection(consulAddr string, serviceName string) (*Client, error) {
	dsn := strings.Join([]string{"consul:/", consulAddr, serviceName}, "/")
	conn, err := grpc.Dial(
		dsn,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		logger.Error("start gRPC connection failed: ", zap.String("err", err.Error()))
	}
	c := pb.NewCommonClient(conn)
	return &Client{c: c, conn: conn}, nil
}

// Close grpc connection closed
func (cli *Client) Close() error {
	return cli.conn.Close()
}

// SendSocket message
func (cli *Client) SendSocket(node string, socketID uint32, ipAddr string, traceID uint64, seqID uint64, cmd uint32, data []byte, timeout uint) (code uint32, rTraceID uint64, rSeqID uint64, rCmd uint32, rData []byte) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancle()
	req := pb.SimpleRequest{Node: node, SocketID: socketID, IpAddr: ipAddr, TraceID: traceID, Cmd: cmd, SeqID: seqID, Data: data}
	reply, err := cli.c.SocketFeature(ctx, &req)
	if err != nil {
		logger.Error("gRPC response error: ", zap.String("err", err.Error()))
		code, rTraceID, rSeqID, rCmd, rData = 1, traceID, seqID, cmd, []byte{}
		return
	}
	code, rTraceID, rSeqID, rCmd, rData = reply.Code, reply.TraceID, reply.SeqID, reply.Cmd, reply.Data
	return
}

// SendJSON message
func (cli *Client) SendJSON(traceID uint64, seqID uint64, path string, data []byte, timeout uint) (rTraceID uint64, rSeqID uint64, rData []byte) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancle()
	req := pb.JsonRequest{TraceID: traceID, SeqID: seqID, Path: path, Data: data}
	reply, err := cli.c.JSONFeature(ctx, &req)
	if err != nil {
		logger.Error("gRPC response error: ", zap.String("err", err.Error()))
		rTraceID, rSeqID, rData = traceID, seqID, sysError
		return
	}
	rTraceID, rSeqID, rData = reply.TraceID, reply.SeqID, reply.Data
	return
}

// func (cli *Client) ServerStreamRequest(ctx context.Context, traceID string, cmd string, seq int32, data []byte) (pb.EMService_ServerStreamCallClient, error) {
// 	req := pb.EMReq{TraceID: traceID, Cmd: cmd, Seq: seq, ReqData: data}
// 	return cli.c.ServerStreamCall(ctx, &req)
// }

// func (cli *Client) GameDataRequest(ctx context.Context, traceID string, cmd string, seq int32, data []byte) (pb.EMService_GameDataCallClient, error) {
// 	req := pb.EMReq{TraceID: traceID, Cmd: cmd, Seq: seq, ReqData: data}
// 	return cli.c.GameDataCall(ctx, &req)
// }
