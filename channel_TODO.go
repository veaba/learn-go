/**
@desc 通道
*/
package main

import (
	"fmt"
	"time"
)

func sum1(a []int, channal chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	channal <- sum // 将sum 送给channel
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6} // len 6
	channel := make(chan int)

	fmt.Println("开始时间：", time.Now())
	go sum1(arr[len(arr)/2:], channel) //后三个，
	go sum1(arr[:len(arr)/2], channel) //前三个，

	fmt.Println("结束时间：", time.Now())

	//sum1(arr[1:],channel)
	//开始时间： 2019-05-21 16:35:23.2390008 +0800 CST m=+0.002000001
	//结束时间： 2019-05-21 16:35:23.2580008 +0800 CST m=+0.021000001
	x, y := <-channel, <-channel //从channel 获取

	fmt.Println(x, "+", y, "=", x+y)

	fmt.Println(arr[len(arr)/2])

	fmt.Println(len(arr))     //6
	fmt.Println(len(arr) / 2) //3
	fmt.Println(arr[3:])      //[4 5 6]

	fmt.Println(arr[:3]) //[1 2 3]

}
