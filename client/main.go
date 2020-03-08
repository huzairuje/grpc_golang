package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"grpc_golang/common/config"
	"grpc_golang/common/model"
	"grpc_golang/response"
	"log"
)

func FirstUser() *model.User {
	return &model.User{
		Id:       "n001",
		Name:     "Ahmad Alpakih",
		Password: "alpakihahmad",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
}

func SecondUser() *model.User {
	return &model.User{
		Id:       "n002",
		Name:     "Muhammad Huzair",
		Password: "wewewe",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
}

func FirstGarage() *model.Garage {
	return &model.Garage{
		Id:   "q001",
		Name: "Garage1",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
}

func SecondGarage() *model.Garage {
	return &model.Garage{
		Id:   "q002",
		Name: "Garage2",
		Coordinate: &model.GarageCoordinate{
			Latitude:  43.123123123,
			Longitude: 52.1231313123,
		},
	}
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

func init()  {
	user := serviceUser()
	fmt.Println("\n", "===========> user test")
	// register user1

	user.Register(context.Background(), FirstUser())
	// register user2

	user.Register(context.Background(), SecondUser())
	garage := serviceGarage()
	fmt.Println("\n", "===========> garage test A")

	garage.Register(context.Background(), FirstGarage())
	garage.Register(context.Background(), SecondGarage())
	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: FirstUser().Id,
		Garage: FirstGarage(),
	})
	// add garage2 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: SecondUser().Id,
		Garage: SecondGarage(),
	})
}

func main() {

	e := echo.New()
	api := e.Group("/api")
	user := serviceUser()
	garage := serviceGarage()
	api.GET("/list/garage", func(c echo.Context) error {
		res3, err := garage.ListAllGarage(context.Background(), new(empty.Empty))
		if err != nil {
			log.Fatal(err.Error())
		}
		if len(res3.List) > 1 {
			return response.List(c, "Success Get List Garage", res3, nil)
		}
		return response.SingleData(c, "Success Get List Garage", res3, nil)
	})

	api.GET("/list/garage/:user_id", func(c echo.Context) error {
		paramUserId := c.Param("user_id")
		res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: paramUserId})
		if err != nil {
			log.Fatal(err.Error())
		}
		if len(res3.List) > 1 {
			return response.List(c, "Success Get List Garage By User", res3, nil)
		}
		return response.SingleData(c, "Success Get List Garage By User", res3, nil)
	})

	api.GET("/list/user", func(c echo.Context) error {
		res1, err := user.List(context.Background(), new(empty.Empty))
		if err != nil {
			log.Fatal(err.Error())
		}
		if len(res1.List) > 1 {
			return response.List(c, "Success Get List Garage By User", res1, nil)
		}
		return response.SingleData(c, "Success Get List Garage By User", res1, nil)
	})

	e.Logger.Fatal(e.Start(":1430"))

}