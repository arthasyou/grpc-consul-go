package consul

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/luobin998877/go_utility/logger"
	"go.uber.org/zap"
)

// Service service agnet of consul
type Service struct {
	id    string
	agent *api.Agent
}

var (
	checkTime      = time.Duration(10) * time.Second
	unregisterTime = time.Duration(1) * time.Minute
)

// RegisterService register the service at consul
func RegisterService(consulAddress string, serviceName string, serviceIP string, servicePort int, tags []string) (*Service, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulAddress
	client, err := api.NewClient(consulConfig)
	if err != nil {
		logger.Error("Consul Register Service error: ", zap.String("err", err.Error()))
		return nil, err
	}

	agent := client.Agent()
	id := fmt.Sprintf("%v-%v-%v", serviceName, serviceIP, servicePort)
	reg := &api.AgentServiceRegistration{
		ID:      id,
		Name:    serviceName,
		Tags:    tags,
		Port:    servicePort,
		Address: serviceIP,
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       checkTime.String(),                                           // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", serviceIP, servicePort, serviceName), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: unregisterTime.String(),                                      // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		logger.Error("Service Register error", zap.String("err", err.Error()))
		return nil, err
	}

	service := &Service{id: reg.ID, agent: agent}

	return service, nil

}

// UnregisterService unregister service from consul
func (service *Service) UnregisterService() error {
	return service.agent.ServiceDeregister(service.id)
}
