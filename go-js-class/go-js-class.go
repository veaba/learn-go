package main

import "fmt"

/** 定义 js class 的成员 */
type Person struct {
	Name string
	Age  int
}

/**
* 构造函数
* @return Person
 */
func newPerson(name string, age int) *Person {
	return &Person{name, age}
}

/** 写一个 person  的方法 */
func (p *Person) Greet() {
	fmt.Println("hello, my name is ", p.Name, '\n')
}

/**
* 构造函数的方法，对引用的值添加 1
 */
func (p *Person) AddAge() {
	p.Age++
	fmt.Println("my age is ", p.Age)
}

func (p *Person) GetInfo() {
	fmt.Println("name ", p.Name, " age ", p.Age)
}

func main() {

	p := newPerson("Tom", 18)
	p.Greet()
	p.GetInfo()
	p.AddAge()
	p.GetInfo()

}
