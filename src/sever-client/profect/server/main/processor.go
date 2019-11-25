package main

import (
	"fmt"
	"io"
	"net"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
	"sever-client/profect/server/process1"
)

type Process2 struct {
	Conn net.Conn
}

//ServerProcessMes 函数 ：根据客户端发送消息的种类不同，决定调用哪个函数
func (this *Process2)ServerProcessMes(mes *message.Message) (err error){
	fmt.Println("mes = ",mes)  // &{SmMes {"content":"LO123","id":1,"pwd":"","name":"","statu":0}}
	switch mes.Type {
	    case message.LoginMesType:
			//处理登录
			up := &process1.UserloginProcess{
				Conn:this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
			up := &process1.UserloginProcess{
				Conn:this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmMesType:
			//处理群发
			up := &process1.UserSM{
				Conn:this.Conn,
			}
			up.ServerProcessSM(mes)
	default:
		fmt.Println("S")
	}
	return
}

func (this *Process2)P() (err error) {
	//循环的客户端发送的信息
	for{
		tf := &method.Transfer{
			Conn: this.Conn,
		}
		//这里读取数据包，直接封装成一个readpkg()
		mes , err := tf.Readpkg()//监听链接
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务端也退出")
				return err
			}else {
				fmt.Println("readpkg err = ",err)
				return  err
			}
		}
		//fmt.Println("mes = ",mes)
		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}