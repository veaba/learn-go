/**
@desc 类型转换
@语法 type_name(表达式)
*/
package main

import "fmt"

func main() {
	//int-float
	intToFloat()
}

func intToFloat() {
	var sum = 999
	var mean float32
	mean = float32(sum)

	mean = float32(sum) / float32(7)
	fmt.Println(mean) //142.71428

	var i = float32(9) / 7
	fmt.Println(i) //1.2857143
}
