package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUnit struct {

}

func (mu *MathUnit) Cart(req float32 , resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

func main()  {
	mathunit := new(MathUnit)
	//调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathunit)
	if err!= nil{
		panic(err.Error())
	}
	//通过该函数把mathunit中提供的服务注册在HTTP协议上，方便调用者利用http的方式进行数据传递
	rpc.HandleHTTP()

	listen , err := net.Listen("tcp",":8881")
	if err != nil{
		panic(err.Error())
	}
	http.Serve(listen,nil)
}