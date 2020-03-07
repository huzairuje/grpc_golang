package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"grpc_golang/common/config"
	"grpc_golang/common/model"
	"log"
	"net"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct{}

func (UsersServer) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)
	log.Println("Starting RPC server at", config.ServiceUserPort)
	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}

	log.Fatal(srv.Serve(l))
}