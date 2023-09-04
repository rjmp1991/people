package main

import (
	"flag"
	"log"

	"github.com/rjmp1991/people/pb"
	"github.com/rjmp1991/people/service"
	"google.golang.org/grpc"
)

func main() {

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("server started on port: %d", *port)
	userServer := service.NewUserServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

}
