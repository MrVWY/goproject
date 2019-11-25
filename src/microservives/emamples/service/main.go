package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"microservives/emamples/service/data"
	O "microservives/emamples/service/proto"
	"time"

)
const path  = "C:/Users/36527/go/src/"
type S struct {

}

func (s *S) Create(ctx context.Context, request *O.Request,response *O.Response)  error   {
	log.Print("Received OrderCreate.Create request")
	response.Order = &O.Order{
		UserId:               request.UserId,
		Name:                 request.Name,
		Price:                request.Price,
		Createtime:           time.Now().Unix(),
		Status:               "Successfully",
	}
	response.Consignmentid = request.Consignmentid
	response.Status = "Successfully"
	//下面错误代码
	//response = &O.Response{
	//	Consignmentid: request.Consignmentid,
	//	Status: "Successfully",
	//	Order:response.Order,
	//}
	//fmt.Println(response)

	//调起支付

	return nil
}

func (s *S) GetAll(ctx context.Context, request *O.Request,response *O.Response)  error   {
	filepath := path + "micro/emamples/data/config.json"
	userid  := request.UserId
	k := data.Readfile(filepath,userid)
	if len(k) == 0 {
		//下面写法错误
		//response = &O.Response{
		//	Status: "fail",
		//	Err:"The userid is not exists",
		//}
		//fmt.Println(response)
		response.Status = "fail"
		response.Err = "The userid is not exists"
		fmt.Println(response)
		return nil
	}
	var o = []*O.Order{}
	for i := 0;i<len(k);i++ {
		t := &O.Order{
			UserId:               k[i].UserId,
			Name:                 k[i].Name,
			Price:                k[i].Price,
		}
		o = append(o,t)  //
	}
	//response = &O.Response{
	//	Status: "Successfully",
	//	Orders:o,
	//}
	//fmt.Println(response)
	response.Status = "Successfully"
	response.Orders = o
	fmt.Println(response)
	return nil
}

func main() {
	service := micro.NewService(
			micro.Name("go.micro.api.CreateOrder"),
	)
	service.Init()
	O.RegisterORDERHandler(service.Server(),new(S))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}