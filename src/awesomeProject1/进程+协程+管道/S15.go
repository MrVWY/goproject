package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int,10)
	//声明一个全局互斥锁
	//lock 是一个全局的互斥锁
	lock sync.Mutex
)

func test1(n int)  {
	res :=1
	for i :=1;i<=n;i++{
		res *=i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main()  {
	for i := 1; i<=20;i++{
		go test1(i)
	}
	time.Sleep(time.Second*10)//休眠10秒
	lock.Lock()
	for i,v := range myMap{
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}