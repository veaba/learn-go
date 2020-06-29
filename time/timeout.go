/**
* @desc 延时2s执行
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1111=>")
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		fmt.Println("i====>", i)
		// time.Sleep(1000)
		time.Sleep(time.Second * 2)
	}

}
