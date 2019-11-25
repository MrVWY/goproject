package main

import "fmt"

func main()  {
	const (
		a = iota
		b
		c = iota
		d,e = iota,iota
	)
	fmt.Println(a,b,c,d,e)
}