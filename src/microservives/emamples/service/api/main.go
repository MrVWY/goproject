package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
	O "microservives/emamples/service/proto"
	B "microservives/emamples/service/model/db"
	Z "microservives/emamples/service/model"
	"strings"
)

type S struct {
	client O.ORDERService

}
//创建订单
func (s *S) CO(ctx context.Context,request *api.Request,response *api.Response) error {
	log.Print("Received Say.CO API request")
	user_id, _ := request.Get["user_id"]
	Consignmentid, _ := request.Get["Consignmentid"]
	name, ok := request.Get["name"]
	price, _ := request.Get["price"]
	db := B.Initialization()
	//根据username查询
	err, _ := B.R(db,&Z.User{
		Username:  strings.Join(name.Values,""),
	})
	if ok && err != nil {
		return errors.BadRequest("go.micro.Usrv.CreateOrder","id cannot exits")
	}
	res, err := s.client.Create(ctx,&O.Request{
		UserId:                   strings.Join(user_id.Values,""),
		Consignmentid:			  strings.Join(Consignmentid.Values,""),
		Name:					  strings.Join(name.Values,""),
		Price:                    strings.Join(price.Values,""),

	})
	if err != nil {
		return err
	}
	fmt.Println(res)
	response.StatusCode = 200
	b, _ := json.Marshal(map[string]*O.Response{
		"response" : res,
	})
	response.Body = string(b)
	return nil
}

func (s *S) SO(ctx context.Context,request *api.Request,response *api.Response) error{
	log.Print("Received Say.SO API request")
	username, ok := request.Get["username"]
	fmt.Println(username)
	db := B.Initialization()
	//根据username查询
	err, _ := B.R(db,&Z.User{
		Username:  strings.Join(username.Values,""),
	})
	if ok && err != nil {
		return errors.BadRequest("go.micro.Usrv.CreateOrder","id cannot exits")
	}
	res, err :=s.client.GetAll(ctx,&O.Request{
		Name:               strings.Join(username.Values,""),
	})
	if err != nil {
		return err
	}
	fmt.Println("res = ",res)
	response.StatusCode = 200
	b, _ := json.Marshal(map[string]*O.Response{
		"response":res,
	})
	response.Body = string(b)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.CreateOrder1"),
		micro.Address(":9000"),
	)
	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&S{client: O.NewORDERService("go.micro.api.CreateOrder",service.Client())},
			),
	)
	//O.RegisterCreateOrderHandler()
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}