package service

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/rjmp1991/people/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (server *UserServiceServer) UpdateUser(ctx context.Context, user *pb.User) (*pb.UserRequest, error) {
	if user.UserId == 0 || user.UserName == "" {
		return nil, fmt.Errorf("%v", codes.InvalidArgument)
	}
	if _, ok := server.Users[user.UserId]; !ok {
		return nil, fmt.Errorf("%v", codes.NotFound)
	}
	server.Users[user.UserId] = user
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

func (server *UserServiceServer) PutUsers(stream pb.UserService_PutUsersServer) error {
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}

		if _, ok := server.Users[req.User.UserId]; ok {
			return logError(status.Errorf(codes.AlreadyExists, "id: %v", req.User.UserId))
		}
		server.Users[req.User.UserId] = req.User
		err = stream.Send(&pb.UserRequest{
			UserId: req.User.UserId,
		})
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
