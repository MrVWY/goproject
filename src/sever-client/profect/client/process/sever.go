package process
import (
	"encoding/json"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
	"fmt"
	"net"
	"os"
)

func ShowMenu()  {
	fmt.Println("------------welcome to this chatroom,sir!------------------")
	fmt.Println("---------1、Display a list of online users---------------------")
	fmt.Println("---------2、Send the message---------------------")
	fmt.Println("---------3、A list of message---------------------")
	fmt.Println("---------4、Exit the System---------------------")
	var key int
	var content string
	smprocess := &Smprocess{}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println("当前在线用户")
			outputonlineUser()
		case 2:
			//fmt.Println("---------2、Send the message---------------------")
			fmt.Println("请输入要发送的消息")
			fmt.Scanf("%s\n",&content)
			smprocess.SendGP(content)
		case 3:
			fmt.Println("---------3、A list of message---------------------")
		case 4:
			fmt.Println("---------4、Exit the System---------------------")
			os.Exit(0)
		default:
			fmt.Println("Your input is incorrect")
	}
}


func severProcessResMes(conn net.Conn)  {
	tf := &method.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("client is waiting the sever send the message")
		mes ,err := tf.Readpkg()
		if err != nil{
			fmt.Println("tf Readpkg err = ",err)
			return
		}
		//如果读取到消息，进行下一步处理逻辑
		switch mes.Type {
			case message.NotifyUserStatusMesType:
				fmt.Println("有人上线了")
				//1 、取出 NotifyUserStatusMes ， 2、把这个用户得消息，状态保存在客户端的map[int]User中
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Date),&notifyUserStatusMes)
				updateUserStatus(&notifyUserStatusMes)
			case message.SmMesType:
				outputSM(&mes)
			default:
				fmt.Println("返回了未知消息类型")
		}
	}
}