package main

import (
	"log"
	"net"

	"github.com/rjmp1991/people/pb"
	"github.com/rjmp1991/people/service"
	"google.golang.org/grpc"
)

func main() {

	userServiceServer := service.NewUserServer()

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
