/**
@desc go没有类
@&x  取变量的指针地址
@*x 取指针指向的内容

- 在结构体上定义方法
- 方法名称出现在 func 和方法名之间的参数中

- 可以对包中的 任意类型类型 定义方法，不仅是结构体
- 但不能对来自其他包的类型和基础类型进行定义方法
*/

package main

import (
	"fmt"
	"math"
)

type Class struct {
	x,y float64
}

/**
-  * 地址

- func (v *Class) Woo() float64  {} 区别
- func (v Class) Woo() float64  {}  区别
*/

func (v Class) Woo() float64  {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

func main()  {
	v :=&Class{3,4}
	fmt.Println(v.Woo())
}
