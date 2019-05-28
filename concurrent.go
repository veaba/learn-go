/**
@desc 并发
@goroutine 轻量级线程
*/
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("我是第一个")
	go say("2019MSI G2获得冠军")
	say("hello")
	fmt.Println("那我呢")
}

// go say("2019MSI G2获得冠军") 与 say("hello") 没有固定顺序，因为它们是两个goroutine 在执行
