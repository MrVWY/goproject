package main
//reflect 反射

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}){
	//通过反射获取传入的变量的type kind 值
	//1、先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = " , rType)

	//2、获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2)
	fmt.Printf("rVal=%v rVal-type=%T \n",rVal,rVal)

	//下面将rVal转成interface{}
	iv := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=",num2)

}

func reflectTest02(b interface{}){
	//通过反射获取传入的变量的type kind 值
	//1、先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = " , rType)

	//2、获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//下面将rVal转成interface{}
	iv := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	fmt.Printf("iv=%v iv-type=%T  \n",rVal,rVal)
	//获取 变量对应的kind   kind是类别  Type是类型
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind = %v  kind = %v \n",kind1,kind2)
	//将 interface{} 通过断言转成需要的类型
	stu,ok := iv.(Student)
	if ok{
		fmt.Printf("stu.name=%v \n",stu.Name)
	}
}

func reflectTest03(b interface{}){
	//获取到reflect.Value
	rVal := reflect.ValueOf(b)
	//rVal 的kind
	fmt.Printf("rVal kind = %v \n",rVal.Kind())
	//rVal
	rVal.Elem().SetInt(20)
}

type Student struct {
	Name string
	Age int
}

func main() {
	var num int = 100
	reflectTest01(num)

	fmt.Println("-----------------------------------")

	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)

	fmt.Println("-----------------------------------")

	var nummber int =10
	reflectTest03(&nummber)
	fmt.Println("num=",nummber)
	/*
	rVal.Elem() ->>>Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	num := 1
	ptr *int = &num
	num := *ptr
	*/
}