package data

import (
	"encoding/json"
	"io/ioutil"
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

//C:/Users/36527/go/src/micro/emamples/service/config.json
func Readfile(filepath string, user_id string) (k []K) {
	//func main() {
	JsonParse := NewJsonStruct()
	var result map[string]interface{}
	JsonParse.Load(filepath, &result)
	node := result["Orders"].([]interface{})
	//fmt.Println(node[0])
	var n = []K{}

	for _, node := range node{
		var m = node.(map[string]interface{})
		//fmt.Println("m = ",m["UserId"])
		if UserId , exists  := m["UserId"] ; exists {
			if UserId == user_id {
				var k K = K{
					UserId:     m["UserId"].(string),
					Name:       m["Name"].(string),
					Price:      m["Price"].(string),
					Createtime: m["Createtime"].(float64),
				}
				n = append(n,k)
			}
		}else {
			return nil
		}
	}
	return n
}
//test
//func main()  {
//	filepath := "C:/Users/36527/go/src/micro/emamples/data/config.json"
//	userid  := "cust0045"
//	k,err := Readfile(filepath,userid)
//	if err != nil{
//		return
//	}
//	fmt.Println(len(k))
//	fmt.Println(k[0])
//	fmt.Println(k[1])
//}