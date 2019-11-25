package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	b:= s
	b[1] = 3
	fmt.Println(b)
	s=s[0:0]
	fmt.Println(s,b)
	fmt.Printf("s: %p b: %p  \n",&s,&b)
	s=append(s,4,5,6)
	fmt.Println(s,b)
}