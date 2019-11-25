package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string,MaxIdle,MaxActive int,IdleTimeout time.Duration)  {
	pool = &redis.Pool{
		MaxIdle:MaxIdle,
		MaxActive:MaxActive,
		IdleTimeout:IdleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",address)
		},
	}
}