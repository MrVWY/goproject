package process

import (
	"encoding/json"
	"fmt"
	"sever-client/profect/common/message"
)

func outputSM(mes *message.Message)  {
	var smrMes message.SmMes
	err := json.Unmarshal([]byte(mes.Date) , &smrMes)
	if err != nil{
		fmt.Println("client outputSM json Unmarshal fail err = ",err)
		return
	}
	//显示
	content := fmt.Sprintf("用户id=\t%d 消息=\t%s",smrMes.Id,smrMes.Content)
	fmt.Println(content)
}