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

var localStorage *model.GarageListByUser
var localStorageRegister *model.GarageList

func init() {
	localStorage = new(model.GarageListByUser)
	localStorageRegister = new(model.GarageList)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct{}

func (s GaragesServer) Register(ctx context.Context, param *model.Garage) (*empty.Empty, error) {
	localStorageRegister.List = append(localStorageRegister.List, param)
	log.Println("Registering Garage", param.String())
	return new(empty.Empty), nil
}

func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage
	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)
	log.Println("Adding garage", garage.String(), "for user", userId)
	return new(empty.Empty), nil
}

func (GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId
	return localStorage.List[userId], nil
}

func (GaragesServer) ListAllGarage(context.Context, *empty.Empty) (*model.GarageList, error) {
	return localStorageRegister, nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)
	log.Println("Starting RPC server at", config.ServiceGaragePort)
	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(srv.Serve(l))
}