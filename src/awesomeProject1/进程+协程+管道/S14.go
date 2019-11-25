package main
// 进程 线程 协程
import (
	"fmt"
	"strconv"
)

func test() {
	for i := 1 ; i<=10; i++{
		fmt.Println("test() Hello test"+strconv.Itoa(i))
		//time.Sleep(time.Second)
	}
}

func main() {
	 go test() //开启一个协程

	for i:=1 ; i<=10; i++{
		fmt.Println("main() Hello main"+strconv.Itoa(i))
	}
}