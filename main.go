package main

import (
	"net"

	"github.com/Serabe/iaas/iaas"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5678")
	if err != nil {
		grpclog.Fatalf("could not open connection %s", err)
	}

	server := grpc.NewServer()

	iaas.RegisterIaasServiceServer(server, iaas.NewIaasServiceServer())
	server.Serve(lis)
}
