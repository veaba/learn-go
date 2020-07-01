/**
@desc go 并发、调度器、goroutine
@goroutine 是轻量级协程，调度由runtime管理

*/
package  main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s,i)
	}
}

func main() {
	go say("world")
	say("hello")
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