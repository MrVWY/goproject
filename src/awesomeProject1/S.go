package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {
	//m :=map[string]string{"type":"10","msg":"hello."}
	//mjson,_ :=json.Marshal(m)
	//mString :=string(mjson)
	//fmt.Printf("print mString:%s",mString)

	message :=  map[string]string{
		"LocalDefaultAddressSpace" : "Config.Message.localAddressSpace",
		"GlobalDefaultAddressSpace":  "Config.Message.globalAddressSpace",
	}
	data, err := json.Marshal(message)
	if  err != nil{
		panic("<GetDefaultAddressSpaces> json.Marshal :" + err.Error())
	}
	messages := base64.StdEncoding.EncodeToString(data)
	fmt.Println(messages)
}