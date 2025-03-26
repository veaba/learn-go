package main

import "fmt"

func main() {
	var a, b, c string

	fmt.Println(a, b, c) // '' ,'' ,''

	var d, e, f int
	var g, h, i float64
	fmt.Println(d, e, f) // 0 0 0
	fmt.Println(g, h, i) // 0 0 0

	var j bool
	fmt.Println(j) // false

	// 短变量声明，至少有一个未被声明6+

	a1, b1 := 1, 2

	fmt.Println(a1, b1)
	a2, b1 := 3, 4
	fmt.Println(a2, b1)

	i1 := 999
	// &i1 指针地址
	// *i1 这个指针地址的值
	fmt.Println(*&i1)

}
