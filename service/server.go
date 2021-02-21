package service

import (
	"context"
	"errors"
	"net"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/arthasyou/grpc-consul-go/consul"
	"github.com/arthasyou/grpc-consul-go/pb"
	"github.com/arthasyou/utility-go/logger"
)

// Service detail
type Service struct {
	consul   *consul.Service
	grpc     *grpc.Server
	health   *HealthImpl
	listener net.Listener
}

// CreateService serive
func CreateService(consulAddr string, port int, name string, tags []string) (*Service, error) {
	serviceAddr := localIP()
	if serviceAddr == "" {
		return nil, errors.New("can't find local IP")
	}
	addr := ":" + strconv.Itoa(port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("Grpc service failed to listen", zap.String("port", addr), zap.String("err", err.Error()))
		return nil, err
	}
	s := grpc.NewServer()
	healthServer := &HealthImpl{}
	pb.RegisterCommonServer(s, &server{})
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	consulService, err := consul.RegisterService(consulAddr, name, serviceAddr, port, tags)
	if err != nil {
		return nil, err
	}

	var service Service
	service.consul = consulService
	service.grpc = s
	service.health = healthServer
	service.listener = lis
	return &service, nil
}

// Start loop
func (s *Service) Start() {
	if err := s.grpc.Serve(s.listener); err != nil {
		logger.Error("GRPC server failed to serve", zap.String("err", err.Error()))
	}
}

// Stop service
func (s *Service) Stop() {
	if s.consul != nil {
		s.consul.UnregisterService()
	}
	// r.healthServer.SetServingStatus(r.regCfg.ServiceName, healthpb.HealthCheckResponse_NOT_SERVING)
	if s.grpc != nil {
		s.grpc.GracefulStop()
	}
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// HealthImpl grpc 健康检查
// https://studygolang.com/articles/18737
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
// https://github.com/hashicorp/consul/blob/master/agent/checks/grpc.go
// consul 检查服务器的健康状态，consul 用 google.golang.org/grpc/health/grpc_health_v1.HealthServer 接口，实现了对 grpc健康检查的支持，所以我们只需要实现先这个接口，consul 就能利用这个接口作健康检查了
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch HealthServer interface 有两个方法
// Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
// Watch(*HealthCheckRequest, Health_WatchServer) error
// 所以在 HealthImpl 结构提不仅要实现 Check 方法，还要实现 Watch 方法
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

type server struct {
	pb.UnimplementedCommonServer
}
