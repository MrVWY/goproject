package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Init() (redis.Conn, error){
	conn, err := redis.Dial("tcp",fmt.Sprintf("%s:%d",Host,Port))
	if err != nil{
		panic(err)
	}
	return conn, nil
}

//func main() {
//	var wg sync.WaitGroup
//	for i:=0 ; i<=5;i++ {
//		wg.Add(1)
//		go Init()
//	}
//	wg.Wait()
//}

func SavenTokenCache(conn redis.Conn)  {

}