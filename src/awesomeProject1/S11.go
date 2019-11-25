package main

import "fmt"

func main()  {
	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@abcd"
	//使用切片获取到abcd
	slice := str[6:]
	fmt.Println("slice =",slice)
	//string是不可变的，也就说不能通过 str[0]='z' 方式修改字符串
	//str[0] = 'z' 编译不会通过 报错 原因是string是不可变
	//要修改字符串 可以先将string -> []byte / 或者  []rune -> 修改 -> 重写转成string
	arr := []byte(str)
	arr[0] =  'z'
	str = string(arr)
	fmt.Println("str = ",str)

	//转成[]byte后 可以处理英文和数字 但是不能处理中文
	//因为[]byte 按字节处理 而一个汉字是三个字节 因此报错
	//解决：把string转成 []rune 即可  rune[]按字符来处理
	arr1 := []rune(str)
	arr1[0] = '我'
	str = string(arr1)
	fmt.Println("str = ",str)
}