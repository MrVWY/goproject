package message1

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main(){
	client , err := grpc.Dial("localhost:8881",grpc.WithInsecure())
	if err != nil{
		panic(err.Error())
	}
	defer client.Close()

	timestamp := time.Now().Unix()
	request := &OrderRequest{OrderId: "201960001",TimeStamp:timestamp}
	orderServiceClient := NewOrderServiceClient(client)
	orderInfo , err := orderServiceClient.GetOrderInfo(context.Background(),request)
	if orderInfo != nil{
		fmt.Println(orderInfo)
	}
}