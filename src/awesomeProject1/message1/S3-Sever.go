package message1

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)
type OrderService struct {
	
}

func (os *OrderService) GetOrderInfo(ctx context.Context, request *OrderRequest) (*OrderInfo, error)   {
	orderMap := map[string]OrderInfo{
		"201960001":{OrderId:"201960001",OrderName:"cloth",OrderStaus:"Paid"},
		"201960002":{OrderId:"201960002",OrderName:"Food",OrderStaus:"Paid"},
		"201960003":{OrderId:"201960003",OrderName:"shoes",OrderStaus:"UnPaid"},
	}
	var response *OrderInfo
	current := time.Now().Unix()
	if (request.TimeStamp > current){
		*response = OrderInfo{OrderId: "",OrderName:"",OrderStaus:"Abnormal order status"}
	}else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
				fmt.Println(result)
				return &result , nil
		}else {
			return nil, errors.New("sever error")
		}
	}
	return  response ,nil
}

func main() {
	server := grpc.NewServer()
	RegisterOrderServiceServer(server,new(OrderService))
	lis, err :=net.Listen("tcp",":8881")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}