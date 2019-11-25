package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
	"microservives/emamples/service3/Usrv/model"
	"microservives/emamples/service3/Usrv/model/db"
	"microservives/emamples/service3/Usrv/proto"
	"time"
)

type R struct {}

func (r *R) Login(ctx context.Context, request *proto.Request, response *proto.Response) error {
	DB := db.Initialization()
	defer DB.Close()
	z := db.R2(DB,&model.User{
		Username:   request.Name,
		Password:   request.Password,
	})
	c := &proto.User{
		Name:                 z.Username,
		Pwd:                  z.Password,
	}
	response.User = c
	return nil
}

func (r *R) Register(ctx context.Context, request *proto.Request, response *proto.Response) error   {
	DB := db.Initialization()
	defer DB.Close()
	_, err := db.R1(DB,request.Name)
	if err != nil{
		return err
	}
	db.C(DB,&model.User{
		Username:   request.Name,
		Password:   request.Password,
		Createtime: time.Now().Unix(),
	})
	response.Status = "Register successfully"
	return nil
}

func main() {
	etcds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Name("go.micro.web.svr.user"),
		micro.Registry(etcds),
	)
	service.Init()
	proto.RegisterUserHandler(service.Server(),new(R))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}