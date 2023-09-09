package service

import (
	"context"
	"fmt"

	"github.com/rjmp1991/people/pb"
	"google.golang.org/grpc/codes"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	Users (map[int32]*pb.User)
}

func NewUserServer() *UserServiceServer {
	return &UserServiceServer{
		Users: make(map[int32]*pb.User),
	}
}

func (server *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if req.UserId == 0 {
		return nil, fmt.Errorf("%v", codes.InvalidArgument)
	}
	user, ok := server.Users[req.UserId]
	if !ok {
		return nil, fmt.Errorf("%v", codes.NotFound)
	}
	return &pb.UserResponse{
		User: user,
	}, nil
}

func (server *UserServiceServer) PutUser(ctx context.Context, user *pb.User) (*pb.UserRequest, error) {
	if user.UserId == 0 || user.UserName == "" {
		return nil, fmt.Errorf("%v", codes.InvalidArgument)
	}
	if _, ok := server.Users[user.UserId]; ok {
		return nil, fmt.Errorf("%v", codes.AlreadyExists)
	}
	server.Users[user.UserId] = user
	return &pb.UserRequest{
		UserId: user.UserId,
	}, nil
}

func (server *UserServiceServer) DelUser(ctx context.Context, user *pb.UserRequest) (*pb.UserRequest, error) {
	if user.UserId == 0 {
		return nil, fmt.Errorf("%v", codes.InvalidArgument)
	}
	if _, ok := server.Users[user.UserId]; !ok {
		return nil, fmt.Errorf("%v", codes.NotFound)
	}
	delete(server.Users, user.UserId)
	return &pb.UserRequest{
		UserId: user.UserId,
	}, nil
}

func (server *UserServiceServer) ListUsers(limit *pb.LimitRequest, stream pb.UserService_ListUsersServer) error {
	size := int32(len(server.Users))
	if limit.MaxResults > size {
		limit.MaxResults = size
	}
	var i int32
	for i = 1; i <= limit.MaxResults; i++ {
		stream.Send(&pb.UserResponse{
			User: server.Users[i],
		})
	}
	return nil
}
