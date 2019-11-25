package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
)

type UserProcess struct {

}

//完成登录校验
func (this *UserProcess) Alogin(userID int , userPwd string) (err error){
	//链接到服务器端
	//fmt.Printf("userID = %s userPwd = %s \n",userID,userPwd)
	conn , err := net.Dial("tcp","localhost:8898")
	if err != nil{
		fmt.Println("net Dial err = ",err)
		return
	}
	defer conn.Close()

	//准备通过conn发送消息给服务器
	//先定义消息主体
	var mes message.Message
	mes.Type = message.LoginMesType  //消息主体的type

	//3、创建一个LoginMes 结构体  即消息主体的Date
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	//将loignMes序列化  即消息主体的Date
	data , err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal err =",err)
		return
	}
	// 把data 赋给 mes.data 字段
	mes.Date = string(data)

	//将 mes 进行序列化  即消息主体
	data ,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json Marshal err =",err)
		return
	}

	var buf [4]byte
	//这个是消息的长度
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	//先发送消息的长度,便于服务器接收
	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}
	//发送消息给服务器
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	//处理服务器返回信息
	tf := &method.Transfer{
		Conn:conn,
	}
	mes , err = tf.Readpkg()
	if err !=nil{
		fmt.Println("Readpkg err = ", err)
	}
	//将mes的Data部分反序列化 成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Date),&loginResMes)
	if loginResMes.Code == 200{
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.Id = userID
		CurUser.Status = message.UserOnline
		//fmt.Println("login success")
		//可以显示当前在线用户列表，遍历
		//fmt.Println("当前在线用户列表")
		for _,v := range loginResMes.UserIds{
			if v == userID {
				continue
			}
			fmt.Println("用户id:\t",v)
			//完成 客户端的 	onlineUsers 完成初始化
			user := &message.User{
				Id:     v,
				Status: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		//监听服务器返回的消息，进行下一步逻辑判断
        go severProcessResMes(conn)
		for {
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}
	return
}

//完成注册校验
func (this *UserProcess) Register(userID int , userPwd string ,userName string) (err error)  {
	//链接到服务器端
	//fmt.Printf("userID = %s userPwd = %s \n",userID,userPwd)
	conn , err := net.Dial("tcp","localhost:8898")
	if err != nil{
		fmt.Println("net Dial err = ",err)
		return
	}
	defer conn.Close()

	//准备通过conn发送消息给服务器
	//先定义消息主体
	var mes message.Message
	mes.Type = message.RegisterMesType  //消息主体的type

	//3、创建一个LoginMes 结构体  即消息主体的Date
	var registerMes message.RegisterMes
	registerMes.User.Id = userID
	registerMes.User.Pwd = userPwd
	registerMes.User.Name = userName

	//将loignMes序列化  即消息主体的Date
	data , err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json Marshal err =",err)
		return
	}
	// 把data 赋给 mes.data 字段
	mes.Date = string(data)

	//将 mes 进行序列化  即消息主体
	data ,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json Marshal err =",err)
		return
	}
	var buf [4]byte
	//这个是消息的长度
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	//先发送消息的长度,便于服务器接收
	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}
	//发送消息给服务器
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	//处理服务器返回信息
	tf := &method.Transfer{
		Conn:conn,
	}

	mes , err = tf.Readpkg() // mes 就是 RegisterResMes
	if err !=nil{
		fmt.Println("Readpkg err = ", err)
	}
	//将mes的Data部分反序列化 成LoginResMes
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Date),&RegisterResMes)
	if RegisterResMes.Code == 200{
		//fmt.Println("login success")
		fmt.Println("Register is success")
		os.Exit(0)
	}
	if RegisterResMes.Code == 100 {
		fmt.Println("用户id已经存在，请重新注册...")
		os.Exit(0)
	}
	return
}