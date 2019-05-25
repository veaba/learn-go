/**
@desc 延迟函数的执行，直到上层函数返回
@desc 自下而上
@desc 后进先出的原则
*/
package main

import "fmt"

func main()  {
	defer fmt.Println("world")//4

	fmt.Println("hello")//1

	defer fmt.Println("world1")//3
	fmt.Println("hello1")//2

	for i:=0;i<10 ;i++  {
		defer fmt.Println(i)
	}
}
