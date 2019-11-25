package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client , err := rpc.DialHTTP("tcp","localhost:8881")
	if err != nil{
		panic(err.Error())
	}
	var req float32 //请求值
	req = 3

	var resp *float32
	//同步
	err = client.Call("MathUnit.Cart",req,&resp)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(*resp)

}