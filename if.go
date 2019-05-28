/**
@desc if 便捷语句，可以条件之前执行一个简单的语句，作用域仅在if范围内
*/
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))

	fmt.Println(math.Pow(3, 2)) //9
}
