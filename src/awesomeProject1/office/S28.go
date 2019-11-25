package main

import "fmt"

type Student struct {
	Name string
}

var b interface{} = Student{
	Name:     "aaa",
}

//func main(){
//	fmt.Println(b)
//	var c  = b.(Student) //首先将 b 指向的数据复制一份，然后转换为 Student 类型赋值给 c。
//	c.Name = "bbb"
//	fmt.Println(b.(Student).Name,c.Name)
//}

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}
//defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
//作为函数参数，则在 defer 定义时就把值传递给 defer，并被缓存起来；
//作为闭包引用的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值
func main() {
	    f := F(5)
	    defer func() {
		       fmt.Println(f())
	    }() //9
	    defer fmt.Println(f()) //7
		defer fmt.Println(f()) //6
	    i := f() // 8
	    fmt.Println("i",i)
}