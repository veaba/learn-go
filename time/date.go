package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	x := fmt.Sprintf("当前时间为： %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// x := fmt.Printf(now.Year())
	fmt.Printf(x)
}
