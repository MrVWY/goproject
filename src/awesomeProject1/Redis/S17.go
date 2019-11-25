package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	//链接到Redis
	conn , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis Dial err = ",err)
		return
	}
	//2、通过go 向redis写入数据
	_ , err = conn.Do("HMSET","username","name","1","name1","JAVA","name2","python")
	if err != nil{
		fmt.Println("conn Do1 err = ",err)
		return
	}
	fmt.Println("---------------------------------------")
	r , err2 :=redis.Values (conn.Do("hmget","username","name","name1","name2"))
	fmt.Println(r)
	if err2 != nil{
		fmt.Println("conn Do err = ",err2)
		return
	}
	for _ , v := range r {
		a := v.([]byte)
		fmt.Printf("%s\n",a)
	}
}