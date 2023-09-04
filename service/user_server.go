package service

import (
	"context"

	"github.com/rjmp1991/people/pb"
	"github.com/rjmp1991/people/sample"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (server *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{
		User: sample.NewUser(),
	}, nil
}
func (server *UserServer) mustEmbedUnimplementedUserServiceServer() {}
