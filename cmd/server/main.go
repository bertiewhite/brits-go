package main

import (
	"fmt"

	// need to change naming here
	innergrpc "github.com/bertiewhite/brits-go/internal/grpc"
	grpc "github.com/bertiewhite/brits-go/pkg/proto"
	"google.golang.org/grpc/reflection"
)

func main() {

	queueSvc := innergrpc.NewQueueService()

	server, err := innergrpc.NewServer()
	if err != nil {
		panic(fmt.Sprintf("The worlds ending!!! %s", err.Error()))
	}
	grpc.RegisterMessageQueueServer(server.GetServer(), queueSvc)
	reflection.Register(server.GetServer())

	server.Serve()

}
