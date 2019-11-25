package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
	"strconv"
	"time"
	O "microservives/emamples/service4/Order/osrv/proto"
	"microservives/emamples/service4/Order/osrv/model/db"
	"microservives/emamples/service4/conf"
)

type R struct {}

func (r *R) Create(ctx context.Context, request *O.Request,response *O.Response)  error   {
	log.Print("Received o.Create request")
	response.Order = &O.Order{
		Name:                 request.Name,
		Price:                request.Price,
		Createtime:           time.Now().Unix(),
	}
	response.Consignmentid = request.Consignmentid
	response.Status = "Successfully"
	response.Payurl = "www.******.com"
	i,_ := strconv.ParseInt(request.Price,10,64)
	dbs := db.Initialization()
	defer dbs.Close()
	db.C1(dbs,&db.Order{
		Username:    request.Name,
		Ordernumber: request.Consignmentid,
		Createtime: response.Order.Createtime,
		Price:      i ,
		Status:      "Successfully",
	})
	return nil
}

func (r *R) GetAll(ctx context.Context, request *O.Request,response *O.Response)  error {
	name  := request.Name
	dbs := db.Initialization()
	defer dbs.Close()
	z, err := db.R1(dbs,name)
	if err != nil{
		return err
	}
	if len(z) == 0 {
		response.Status = "not Order "
		return nil
	}else {
		var o = []*O.Order{}
		for i := 0;i<len(z);i++ {
			t := &O.Order{
				Name:                 z[i].Username,
				//Price:                z[i].Price,
				Status:               z[i].Status,
				Createtime:           z[i].Createtime,
			}
			o = append(o,t)
		}
		fmt.Println(z)
		response.Status = "Successfully"
		response.Orders = o
		fmt.Println(response)
		return nil
	}
	return nil
}

func main() {
	etcds := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{conf.Etcd}
	})
	service := micro.NewService(
		micro.Name("go.micro.web.srv.o"),
		micro.Registry(etcds),
	)
	service.Init()
	O.RegisterOHandler(service.Server(),new(R))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}