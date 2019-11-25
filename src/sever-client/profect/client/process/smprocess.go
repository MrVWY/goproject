package process

import (
	"encoding/json"
	"fmt"
	"sever-client/profect/common/message"
	"sever-client/profect/method"
)

type Smprocess struct {}

//群发消息
func (this *Smprocess) SendGP(content string) (err error) {
	//创建message
	var mes message.Message
	mes.Type = message.SmMesType

	var sms message.SmMes
	sms.Content = content
	sms.Id = CurUser.Id
	sms.Status = CurUser.Status

	//序列化
	data , err := json.Marshal(sms)
	if err != nil {
		fmt.Println("SendGP Json1 Marshal err = ",err.Error())
		return
	}

	mes.Date = string(data)
	data , err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGP Json2 Marshal err = ",err.Error())
		return
	}
	//发送给服务器
	tf := &method.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.Writepkg(data)
	if err != nil {
		fmt.Println("SendGP Writepkg  err = ",err.Error())
		return
	}
	return
}