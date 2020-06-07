/**
@channel=> 通信机制，一个goroutine 可以向另外一个goroutine发送消息
@int 发送int类型消息的channel => chan int
@make channel 使用make函数创建
@close 已读取的消息会被清零
@channel type channel 的类型分为=> 带缓存的channel、不带缓存的channel
@noCache channel
	- 读取消息会被阻塞，直到goroutine向该channel中发送消息
	- 向无缓存的channel中发送消息也会被阻塞，直到有goroutine向channel中读取消息
@hasCache channel
	- 声明方式为指定make函数的第二个参数，
	- 该参数为channel缓存的额容量
	- ch := make(chan int,10)
	- 类似阻塞队列，由环形数组实现，
	- 缓存未满，向channel中发送消息时不会阻塞
	- 缓存满时，发送操作将被阻塞
	- channel为空时，读取操作会造成阻塞，直到goroutine 向channel中写入消息
*/

package main

import (
	"fmt"
)

func main() {
	//hasCacheChannel()
	withGoroutine()
}

// close 未被读取的消息将存在
func willBeHas() {
	// channel 创建
	channel := make(chan int, 10)
	fmt.Println(channel) // 0xc00004a060

	// channel 写

	channel <- 11

	close(channel)
	t, ok := <-channel
	fmt.Println(t, ok) //11 true
}

// close已读取的消息会被清零
func willBeClearZero() {
	// channel 创建
	channel := make(chan int, 10)
	fmt.Println(channel) // 0xc00004a060

	// channel 写

	channel <- 11

	close(channel)

	fmt.Println(channel)
	fmt.Println(len(channel))
	for x := range channel {
		fmt.Println(x)
	}
	t, ok := <-channel
	fmt.Println(t, ok) //0 false
}

// 无缓存的channelF
func noCacheChannel() {

}

// 缓存channel
func hasCacheChannel() {
	ch := make(chan int, 3)
	//<-ch // error 此时ch是空的，从空缓存中读取，将会被阻塞
	ch <- 1
	//ch <- 2
	ch <- 5
	fmt.Println(ch)
	fmt.Println(len(ch)) // len得到元素个数
	fmt.Println(cap(ch)) // cap函数得到数组容量
}

// goroutine 与 channel通信
func withGoroutine() {

}
