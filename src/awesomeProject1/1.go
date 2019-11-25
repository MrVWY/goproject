package main

import (
	"fmt"

)

type Visitor struct {
	name string
	age int
}

func (visitor *Visitor) showPrice() {
	if visitor.age > 18{
		fmt.Printf("游客的名字为 %v  年龄为  %v   收费20元 \n",visitor.name,visitor.age)
	}else {
		fmt.Printf("游客的名字为 %v  年龄为  %v   收费20元 \n",visitor.name,visitor.age)
	}
}


//fmt是I/O输出包
func main() {
	//这个是主函数
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")
	var v Visitor
	for {
		fmt.Printf("请输入你的名字")
		fmt.Scanln(&v.name)
		if v.name == "n"{
			fmt.Printf("退出程序........")
			break
		}
		fmt.Printf("请输入你的年龄")
		fmt.Scanln(&v.age)
		v.showPrice()
	}
}
