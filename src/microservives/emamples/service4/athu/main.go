package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
	jwt "microservives/emamples/service3/athu/model/authentication"
	O "microservives/emamples/service3/athu/proto"
	"time"
)
type R struct {}

func (r *R) MakeToken(ctx context.Context,request *O.Request,response *O.Response)  error {
	U := request.Username
	d := time.Duration(1000)*time.Second
	token, err := jwt.JwtGenerateToken(U,d)
	if err != nil{
		response.Error = "The product token is fail"
		return nil
	}
	response.Token =  token
	return nil
}

func (r *R) DelToken(ctx context.Context,request *O.Request,response *O.Response)  error {

	return nil
}

func (r *R) VerificationToken(ctx context.Context, request *O.Request, response *O.Response) error{
	_, err := jwt.ParseToken(request.Token)
	if err != nil {
		log.Println(err)
	}
	response.Error = "Token is express"
	return nil
}

func main() {
	ectds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Name("go.micro.web.srv.athu"),
		micro.Registry(ectds),
	)
	service.Init()
	O.RegisterAthuHandler(service.Server(),new(R))
	if err := service.Run() ; err != nil{
		log.Fatal(err)
	}
}