/**
@desc go 并发、调度器、goroutine
@goroutine 是轻量级协程，调度由runtime管理

常见用于分离定时器和web server 的主程流程
如以下例子:

	fmt.Println("Hello world!~")
	app := iris.New()
	app.Handle(GET, "/", indexPage)
	go timeInterval()
	_ = app.Run(iris.Addr(":8080"))

*/
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		go hello(s, i)
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now(), s, i)
	}
}

func main() {
	go say("go线程")
	say("go 主线程")
}
func hello(h string, i int) {
	fmt.Println("I am hello", h, i)
}

/**
world 0
hello 0
hello 1
world 1
world 2
hello 2
hello 3
world 3
world 4
hello 4
____________________
world 0
hello 0
world 1
hello 1
hello 2
world 2
hello 3
world 3
hello 4


*/
