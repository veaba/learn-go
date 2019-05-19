/**

- 命名变量
- 短函数中，x,y int 可以需要return 名称，会直接返回当前结果的当前值，也就是变量的值 x 和 y的值，定义函数时候的return
*/
package main

import "fmt"

func split(sum int)(x1,y1 int)  {
	x1 = sum *4 /9
	y1 = sum-4

	fmt.Println(sum) // 17
	fmt.Println(x1) //7
	fmt.Println(y1) //13
	return
}
func main()  {
	fmt.Println(split(17)) //7 13
}
