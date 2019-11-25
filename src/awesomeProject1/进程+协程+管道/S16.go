package main

import "fmt"

type  Cat struct {
	Name string
	Age int
}
func main() {
	//创建管道
	var intChan chan  int
	intChan = make(chan int,3)//最大只能放3个
	//
	fmt.Printf("intChan = %v  intChan本身的值 = %v\n ",intChan,&intChan)
	//像管道写入数据
	intChan<-10
	num := 222
	intChan<-num
	fmt.Printf("channne len = %v cap=%v\n",len(intChan),cap(intChan))
	//从管道取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2 = %v\n",num2)
	fmt.Printf("channne len = %v cap=%v\n",len(intChan),cap(intChan))
	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告deadlock

	//列子
	//定义一个存放任意数据类型的管道
	//var allchan chan interface{}
	allchan := make(chan interface{},3)
	allchan<- 10
	cat :=Cat{
		Name: "AAAA",
		Age:  100,
	}
	allchan<-cat
	allchan<-"jack"
	<-allchan
	newcat := <-allchan
	fmt.Printf("newcat=%T,newcat=%v",newcat,newcat)
	//类型断言
	a :=newcat.(Cat)
	fmt.Printf("newcat.Name=%v",a.Name)
}