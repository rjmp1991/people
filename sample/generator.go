package sample

import (
	"math/rand"

	"github.com/rjmp1991/people/pb"
)

func randomId() int32 {
	return int32(rand.Intn(10)) + 1
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomName() string {
	return randomStringFromSet(
		"Jonh Doe",
		"Maria Lee",
		"Juan Lopez",
		"Pepito Perez",
	)
}

// returns a new user
func NewUser() *pb.User {
	user := &pb.User{
		UserId:   randomId(),
		UserName: randomName(),
	}
	return user
}
