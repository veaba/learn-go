/**
@desc 通道缓存区
@cache 带缓存，允许发送和接收的数据处于异步状态
	- 发送的数据可以放到缓冲区
	- 可以等待接收端去获取，而非立刻需要接收端去获取数据
	- 缓冲区满了，发送端将无法再发送数据
	- 发送方会阻塞直到发送的值被拷贝到缓冲区
	- 读出缓存完，就不能再读取了，报错

@noCache 不带缓冲
	- 发送方会阻塞直到接收方从通道中接收值

*/
package main

import "fmt"

func main() {
	cacheChannel()
}

func cacheChannel() {
	ch := make(chan int, 2)

	// ch是带缓冲的通道，可以同时发送两个数据
	// 而不用立刻去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2

	//fmt.Println(<-ch) // 报错了
}
