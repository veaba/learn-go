/**
@desc1 将所有具有共性的方法都放在一起
@desc2 任何其他类型只要实现了这个方法就是实现了这个接口
@desc3 一组方法定义的集合
@这个部分很重要，但不是特别懂，需要和struct一起配合
@TODO 待深入学习
*/
package main

import "fmt"

func main() {
	m := map[string]interface{}{"name": "娃娃。", "age": 999} //和map 格式一样
	fmt.Println(m)

	var phone Phone
	phone = new(OnePlus)
	phone.Call()
	phone.Who()

	phone = new(IPhone)
	phone.Call()
	phone.Who()

	var typeOne OnePlus
	typeOne.name = "全新英雄联盟赛事官网"

	fmt.Println(typeOne)
}

// 定义一个手机的接口
type Phone interface {
	Call()
	Who() //定义了一个方法call()

}

type OnePlus struct {
	name string
}

type IPhone struct {
}

func (onePlus OnePlus) Call() {
	fmt.Println("onePlus Call police")
}
func (onePlus OnePlus) Who() {
	fmt.Println("who who ")
}

func (IPhone IPhone) Call() {
	fmt.Println("iPhone call police")
}
func (IPhone IPhone) Who() {
	fmt.Println("implement me")
}
