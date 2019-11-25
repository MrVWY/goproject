package main
//-------------------------defer语句----------------
import (
	"fmt"
)

func sum(n1 int,n2 int) int {
	//当执行到defer时，暂时不执行，回讲defer后面的语句压入独立的栈中
	defer fmt.Println("ok1 n1 =",n1)
	defer fmt.Println("ok1 n2 =",n2)
	res := n1 + n2
	fmt.Println("ok1 res =",res)
	return res
}

func sum1(n1 int,n2 int) int {
	//当执行到defer时，暂时不执行，回讲defer后面的语句压入独立的栈中
	defer fmt.Println("ok1 n1 =",n1)
	defer fmt.Println("ok1 n2 =",n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok1 res =",res)
	return res
}

func main() {
	sum(20,10)
	fmt.Println("--------------------------------------")
	sum1(20,10)
}