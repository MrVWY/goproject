package 断言
//断言

import "fmt"

type Point struct {
	x int
	y int
}

func main(){
	var a interface{}
	var point Point = Point{1,2}
	a = point  //ok
	var b Point
	//b = a 不可以
	b = a.(Point)//类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错
	fmt.Println(b)

	fmt.Println("-----------------------------------------------------------------")

	var x interface{}
	var b2 float32 =1.1
	x =b2 //空接口 可以接收任何类型
	//x=>float32 [使用类型断言]
	y := x.(float32)
	fmt.Printf("y的类型是 %T ，值是=%v \n",y,y)

	fmt.Println("-----------------------------------------------------------------")
	//断言-带检测
	var x1 interface{}
	var b3 float32 =1.1
	x =b3
	y1,ok := x1.(float64)
	if ok{
		fmt.Println("success")
		fmt.Printf("y的类型是 %T ，值是=%v",y1,y1)
	}else {
		fmt.Println("fail")
	}
	fmt.Println("continue")

}