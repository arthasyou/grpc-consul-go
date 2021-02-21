package main

import (
	"fmt"

	"github.com/luobin998877/go_grpc_with_consul/service"
)

func main() {
	c, err := service.CreateConnection("127.0.0.1:8500", "testS")
	if err != nil {
		return
	}
	code, _, _, cmd, reply := c.SendSocket("shhig", 1, "127.0.0.1", 1, 1, 12344, []byte{1, 2, 3, 4, 5}, 5)
	fmt.Println("code: ", code, "cmd: ", cmd, "reply: ", reply)

	_, _, r := c.SendJSON(1, 1, "api/ihe/ihig", []byte{123, 34, 105, 100, 34, 58, 49, 125}, 5)
	fmt.Println("reply: ", r)
}
