package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"log"
	"microservives/emamples/service4/Order/api/handler"
	"microservives/emamples/service4/conf"
)

func main() {
	etcds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{conf.Etcd}
	})

	service := web.NewService(
		web.Name("go.micro.web.o"),
		web.Registry(etcds),
	)
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		})); err != nil {
		log.Fatal("Init", err)
	}
	service.HandleFunc("/o/Creates", handler.Creates)
	service.HandleFunc("/o/GetAlls",handler.GetAlls)
	if err := service.Run() ; err != nil{
		log.Fatal(err)
	}
}