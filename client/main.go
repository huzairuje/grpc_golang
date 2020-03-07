package main

import (
	"context"
	//"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"grpc_golang/common/config"
	"grpc_golang/common/model"
	"log"
	"net/http"
)

var user1 = model.User{
	Id:       "n001",
	Name:     "Ahmad Alpakih",
	Password: "kw8d hl12/3m,a",
	Gender:   model.UserGender(model.UserGender_value["MALE"]),
}

var user2 = model.User{
	Id:       "n002",
	Name:     "Muhammad Huzair",
	Password: "wewewe",
	Gender:   model.UserGender(model.UserGender_value["MALE"]),
}

var garage1 = model.Garage{
	Id:   "q001",
	Name: "Garage1",
	Coordinate: &model.GarageCoordinate{
		Latitude:  45.123123123,
		Longitude: 54.1231313123,
	},
}

var garage2 = model.Garage{
	Id:   "q002",
	Name: "Garage2",
	Coordinate: &model.GarageCoordinate{
		Latitude:  43.123123123,
		Longitude: 52.1231313123,
	},
}

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {

	e := echo.New()
	api := e.Group("/api")

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	go user.Register(context.Background(), &user1)

	// register user2
	go user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}


	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	// add garage1 to user1
	go garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// add garage2 to user1
	go garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}


	api.GET("/list/garage", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res3)
	})

	api.GET("/list/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res1)
	})

	e.Logger.Fatal(e.Start(":1430"))

}