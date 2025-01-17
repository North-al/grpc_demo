package main

import (
	"fmt"

	"github.com/North-al/grpc_demo/internal/service"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "North",
		Age:      20,
	}

	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	newUser := &service.User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())

}
