package 断言
//断言实例2
import "fmt"

func Typejudge(items...interface{})  {

	for index, x := range items{
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v \n",index,x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v \n",index,x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v \n",index,x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 int 类型，值是%v \n",index,x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v \n",index,x)
		default:
			fmt.Printf("第%v个参数不确定，值是%v \n",index,x)
		}
	}
}

func main(){
	var n1 float64 =1.3
	Typejudge(n1)
}