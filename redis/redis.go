package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main()  {
	client :=redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"yourPassword",
		DB:1,
	})
	pong,err:=client.Ping().Result()
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println(pong)

	v,err1:=client.Do("get","hello").String()
	if err1!=nil{
		return
	}
	fmt.Println("v:",v)
}


