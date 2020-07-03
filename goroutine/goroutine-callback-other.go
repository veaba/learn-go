/**
* @desc 下面的代码无法按预期执行a
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("before")
	go one()
	fmt.Println("after")
	two()
}

func one() {
	fmt.Println("one====>")
	oneSchool()
}

func oneSchool() {
	fmt.Println("one child")
}

func two()  {
	fmt.Println("I am two")
}
