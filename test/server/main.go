package main

import (
	"fmt"

	_ "github.com/luobin998877/go_grpc_with_consul/consul" // very important
	"github.com/luobin998877/go_grpc_with_consul/service"
)

func main() {
	service.Register(&handler{})
	s, _ := service.CreateService("127.0.0.1:8500", 12345, "testS", []string{"ighigh"})
	fmt.Println(s)
	s.Start()
}

type handler struct{}

func (h *handler) HandleCmd(node string, socketID uint32, ipAddr string, cmd uint32, data []byte) (code uint32, rData []byte) {
	fmt.Println("ip: ", ipAddr)
	code = 0
	rData = []byte{1, 2, 3, 4, 5}
	return
}

func (h *handler) HandleJSON(path string, data []byte) (rData []byte) {
	rData = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 44, 32, 34, 109, 115, 103, 34, 58, 34, 115, 121, 115, 32, 101, 114, 114, 111, 114, 34, 125}
	return
}
