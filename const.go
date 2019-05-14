package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	/// const b = "我爱中国" //隐式声明
	// const a string = "I love china" //显示声明

	// fmt.Println(a,b)

	// 常做枚举
	const (
		name = "Jo gel"
		job = "teacher"
		age = 999
	)

	// fmt.Println(name,job,age)

	fmt.Println(unsafe.Sizeof(age))
	fmt.Println(len(name))
	fmt.Println(len(job))
}

