/**
@desc1 将所有具有共性的方法都放在一起
@desc2 任何其他类型只要实现了这个方法就是实现了这个接口
@
*/
package main

import "fmt"

func main()  {
	m:=map[string]interface{}{"name":"娃娃。","age":999} //和map 格式一样
	fmt.Println(m)
}



// 定义一个手机的接口
type Phone interface {
	call()
}

type OnePlus struct {

}

func (onePlus OnePlus)  {

}
