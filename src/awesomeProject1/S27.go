package main

import "fmt"

var m *single
type single struct{
	Name string
}
func GetInstance()*single{
	if m == nil{
		m = &single{}
	}
	return m
}
func main(){
	a := GetInstance()
	a.Name = "a"
	b := GetInstance()
	b.Name = "b"
	fmt.Println(&a.Name, a)
	fmt.Println(&b.Name, b)
	fmt.Printf("%p %T\n", a, a)
	fmt.Printf("%p %T\n", b, b)
}