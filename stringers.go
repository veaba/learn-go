/**
@desc Stringer 是一个可以用字符串串描述自己的类型
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

//自我描述~~~看起来没有调用这个函数，但还是被调用了~
func (p Person) String() string {
	fmt.Println(p.Name)
	return fmt.Sprintf("%v 22 (%v years) \n", p.Name, p.Age)
}

func main() {
	a := Person{"男人", 966, ""}
	b := Person{"女人", 333, ""}
	fmt.Println(a, b)
}
