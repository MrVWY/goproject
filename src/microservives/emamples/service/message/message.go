package message
//                       message/message.go     --Push message!!!!
import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"log"
	"time"
)

var (
	topic = "go.micro.web.topic.order"
	returnTopic = "go.micro.web.topic.pay"
)

func Puborder(name string,price int64,consignmentName string)  {
	msg := &broker.Message{
		Header: map[string]string{
			"time":fmt.Sprintf("%d",time.Now().Unix()),
		},
		Body:[]byte(fmt.Sprintf("%s%s%d",name,consignmentName,price)),
	}
	if err := broker.Publish(topic,msg); err != nil{
		log.Print("[pub] send a fail message",err)
	}else {
		log.Print("[pub] send a message",string(msg.Body))
	}
}

func Pubpay(url string)  {
	msg := &broker.Message{
		Header: nil,
		Body:  []byte(fmt.Sprintf("%s",url)),
	}
	if err := broker.Publish(returnTopic,msg); err != nil{
		log.Print("[pub] send a fail message",err)
	}else {
		log.Print("[pub] send a fail message",string(msg.Body))
	}
}