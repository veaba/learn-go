package main

import "fmt"

/**
- go 与c不同，go没有指针运算
- 间接引用或 非直接引用










*/
func main() {
	i, j := 242, 9999
	fmt.Println(i, j)

	//
	p := &i         //指向i
	fmt.Println(p)  //0xc000060080
	fmt.Println(*p) //242 //通过指针读取i的值
	fmt.Println(&p) //242 0xc00008c020

	*p = 26         //通过指针设置i
	fmt.Println(*p) //26
	fmt.Println(i)  //26

	p = &j //指向j

	fmt.Println()
	*p = *p / 99
	fmt.Println(j)

}
