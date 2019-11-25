package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"log"
	"microservives/emamples/service3/api/handler"
)

func main() {
	etcds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})

	service := web.NewService(
		web.Name("go.micro.web.user"),
		web.Registry(etcds),
	)
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		})); err != nil {
		log.Fatal("Init", err)
	}
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/register",handler.Register)
	if err := service.Run() ; err != nil{
		log.Fatal(err)
	}
}
