package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	O "microservives/emamples/service4/Inventory/proto"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"microservives/emamples/service4/conf"
)

type I struct {}

func (i *I) Shell(ctx context.Context, request *O.Request,response *O.Response) error  {
	orderid := request.BookId
	ordername := request.Bookname
	orderprice := request.Bookprice

	return nil
}

func (i *I) Confim(ctx context.Context, request *O.Request,response *O.Response) error  {
	return nil
}

func main()  {
	ectds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{conf.Etcd}
	})
	service := micro.NewService(
		micro.Name("go.micro.srv.inventorys"),
		micro.Registry(ectds),
	)
	service.Init()
	O.RegisterIHandler(service.Server(),new(I))
	if err := service.Run() ; err != nil{
		log.Fatal(err)
	}
}