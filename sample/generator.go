package sample

import (
	"math/rand"

	"github.com/rjmp1991/people/pb"
)

func randomId() int32 {
	return int32(rand.Intn(10))
}

func randomName() string {
	return "Jonh Doe"
}

// returns a new user
func NewUser() *pb.User {
	user := &pb.User{
		Id:   randomId(),
		Name: randomName(),
	}
	return user
}
