/**
@desc channel 通道遍历和close
@close 如果通道接收不到数据后ok就为false，这时通道可以使用close函数关闭
	- v,ok := <-ch
*/
package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	fmt.Println("n=>", n)
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Println("x=>", i)
	}
	//close(c)
}

func ergodic() {
	n := 10
	for i := 0; i < n; i++ {
		print(i, "\n")
		time.Sleep(1)
	}
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// range 函数遍历每个从通道接收到的数据
	// c 在发送完10个数据后，通道就关闭了
	// so, 这里range函数在接收到10个数据之后就结束
	// 如果上面的c通道不关闭，那么range 函数就不会结束
	// 从而在接收第11个数据的时候就阻塞了
	for i := range c {
		fmt.Println("i=>", i)
	}
	ergodic()
}
