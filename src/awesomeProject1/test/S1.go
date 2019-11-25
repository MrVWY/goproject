package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
//定义配置文件解析后的结构
type K struct {
	UserId string
	Name string
	Price string
	Createtime float64
}

type JsonStruct struct {}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{})  {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

//C:/Users/36527/go/src/shippy/emamples/service/config.json
func Readfile(filepath string, user_id string) {//(k []K) {
	//func main() {
	JsonParse := NewJsonStruct()
	var result map[string]interface{}
	JsonParse.Load(filepath, &result)
	node := result["Orders"].([]interface{})
	//var n = []K{}
	for _, node := range node{
		var m = node.(map[string]interface{})
		fmt.Println("m = ",m["UserId"])
		//if UserId , exists  := m["UserId"] ; exists {
		//	if UserId == user_id {
		//		var k K = K{
		//			UserId:     m["UserId"].(string),
		//			Name:       m["Name"].(string),
		//			Price:      m["Price"].(string),
		//			Createtime: m["Createtime"].(float64),
		//		}
		//		n = append(n,k)
		//	}
		//}else {
		//	return nil
		//}
		var s K = K{
			UserId:     "cust006",
			Name:       "user005",
			Price:     "Manchester, United Kingdom",
			Createtime: 201910282,
		}
		rest := map[string]interface{}{
			"UserId":     "cust006",
			"Name":       "user005",
			"Price":     "Manchester, United Kingdom",
			"Createtime": 201910282,
		}
		m = append(m,rest)
		a ,_ := json.Marshal(m)
		fp ,err := os.OpenFile("C:/Users/36527/go/src/awesomeProject1/test/config.json",os.O_RDWR|os.O_CREATE,0755)
		if err !=nil {
			log.Print("err ",err)
		}
		fmt.Println(fp)
		defer fp.Close()
		fp.Seek(0,2)
		_ , err = fp.Write(a)
		if err !=nil {
			log.Print("err ",err)
		}
		fmt.Println(a)
	}
	//return n
}
//test
func main()  {
	filepath := "C:/Users/36527/go/src/awesomeProject1/test/config.json"
	userid  := "cust003"
	Readfile(filepath,userid)
	//fmt.Println(len(k))
	//fmt.Println(k[0])
	//fmt.Println(k[1])
}