package service

import (
	"context"

	"github.com/rjmp1991/people/pb"
	"github.com/rjmp1991/people/sample"
)

type UserServiceServer struct {
}

func NewUserServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (server *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		User: sample.NewUser(),
	}, nil
}

func (server *UserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
