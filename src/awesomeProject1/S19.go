package main
//Context
import (
	"context"
	"fmt"
	"time"
)

func someHandler() {
	// 创建继承Background的子节点Context
	ctx, cancel := context.WithCancel(context.Background())
	go doSth1(ctx)

	//模拟程序运行 - Sleep 5秒
	time.Sleep(5 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func doSth1(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}

func main() {
	fmt.Println("start...")
	someHandler()
	fmt.Println("end.")
}
