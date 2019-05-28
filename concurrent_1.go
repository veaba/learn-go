/**


rand.Int31()?? 1462604538 1593108510

fmt.Sprintf

time.Second

创建字符串类型的通道

fmt.Println(time.Microsecond) //1µs

fmt.Println(time.Hour) //1h0m0s

fmt.Println(time.Second) //1s


*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// data product channel chan<-string
func producter(header string) {
	// 无限循环，不停生成数据

	//fmt.Println(fmt.Sprintf("%s:%v",header,rand.Int31()))

	//channel<- fmt.Sprintf("%s:%v",header,rand.Int31())

	for {
		// 将随机数和字符串格式化为发送给通道
		//channel<- fmt.Sprintf("%s:%v",header,rand.Int31())
		fmt.Println(fmt.Sprintf("%s:%v", header, rand.Int31()))

		// 等待1s
		time.Sleep(time.Second) //类似js中的 setTimeout
		//time.Sleep(100*time.Microsecond)
	}
}

// data customer

func customer(channel <-chan string) {
	//不停的捕获数据
	for {
		//从通道中取出数据，此处会阻塞直到信道中返回数据
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	//创建一个字符串类型的通道
	//channel := make(chan string)

	//并发 生产者
	//go producter("狗狗：",channel)
	producter("狗狗")

	//go producter("猫猫：",channel)

	//go customer(channel)
}
