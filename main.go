package main

import (
	"fmt"
	"net"

	pb "github.com/menothe/ipg/proto"
	srv "github.com/menothe/ipg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := srv.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterImageProcessorServer(grpcServer, server)

	reflection.Register(grpcServer)

	fmt.Println("Server listening on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}
