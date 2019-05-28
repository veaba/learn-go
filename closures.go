/**
@desc 闭包
- 闭包是一个函数值，来自函数体外部的变量的引用
- 函数可以对该引用值进行访问和赋值
- 这个函数被“绑定”在这个变量上
*/
package main

import "fmt"

func main() {
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(pos(i), neg(-2*i))
	//}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/**
@desc F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*）
- 0 和 1 开始，之后的斐波那契数列系数就由之前的两数相加。
*/
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		temp := x
		x, y = y, (x + y)
		return temp
	}
}

//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
