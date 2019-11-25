package process1

import (
	"encoding/json"
	"fmt"
	"net"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
)

type UserSM struct {
	Conn net.Conn
}

func (this *UserSM) ServerProcessSM(mes *message.Message) {
	//先从mes 取出mes.Data  并直接反序列化成RegisterMes
	var smrMes message.SmMes
	err := json.Unmarshal([]byte(mes.Date) , &smrMes)
	if err != nil{
		fmt.Println("ServerProcessSM json Unmarshal fail err = ",err)
		return
	}
	fmt.Println("ServerProcessSM mes1 = ",mes)
	data , err := json.Marshal(mes)
	if err != nil{
		fmt.Println("ServerProcessSM json Unmarshal fail err = ",err)
		return
	}
	for id , up := range userMgr.onlineUsers{
		if id == smrMes.Id{
			continue
		}
		this.SendElineUser(data,up.Conn)
	}
}

func (this *UserSM) SendElineUser(content []byte , conn net.Conn) {
	tf := &method.Transfer{
		Conn:conn,
	}
	err := tf.Writepkg(content)
	if err != nil {
		fmt.Println("转发消息失败")
	}
}