/**
@desc channel 用来传递数据的一个结构
@goroutine 用于两个goroutine 之间通过传递一个指定类型值来同步运行和通讯
*/
package main

import "fmt"

func sum(s []int, c chan int) {
	fmt.Println("s=>", s)
	total := 0
	for _, v := range s {
		total += v
	}
	fmt.Println("total=>", total)
	c <- total // 把sum发送给通道c
}

func main() {
	s := [] int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s)      // => [7 2 8 -9 4 0]
	fmt.Println(c)      // => 0xc00004a0c0
	fmt.Println(len(c)) // => 0

	go sum(s[:len(s)/2], c) // 后两个
	go sum(s[len(s)/2:], c) // 前两个

	x, y := <-c, <-c // 从通道c中接收数据给x,y
	fmt.Println(x, y, x+y)
}
