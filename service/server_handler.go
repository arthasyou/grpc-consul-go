package service

import (
	"context"

	"github.com/arthasyou/grpc-consul-go/pb"
)

// Handler creates a service handler that will be used to handle message.
type Handler interface {
	HandleCmd(node string, socketID uint32, ipAddr string, cmd uint32, data []byte) (code uint32, reply []byte)
	HandleJSON(path string, data []byte) (relpy []byte)
}

// SocketFeature implements socket data
func (s *server) SocketFeature(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleReply, error) {
	code, reply := handler.HandleCmd(in.Node, in.SocketID, in.IpAddr, in.Cmd, in.Data)
	return &pb.SimpleReply{Code: code, TraceID: in.TraceID, SeqID: in.SeqID, Cmd: in.Cmd, Data: reply}, nil

}

// JsonFeature implements json data
func (s *server) JSONFeature(ctx context.Context, in *pb.JsonRequest) (*pb.JsonReply, error) {
	reply := handler.HandleJSON(in.Path, in.Data)
	return &pb.JsonReply{TraceID: in.TraceID, SeqID: in.SeqID, Data: reply}, nil
}

var handler Handler

// Register with the same name, the one registered last will take effect.
func Register(h Handler) {
	handler = h
}
