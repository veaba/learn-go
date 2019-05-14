/**
@desc go 逻辑运算

*/
package main

import "fmt"

func main() {
	var (
		a1 = 1
		a2 = "hello "
		a3 = "world"
		a4 = 9
		//a4=false
		//a5=true
		//a6="w"
	)
	var b1 = a2 + a3
	var b2 = a1 + a4
	a4+=6
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(a4)
}
