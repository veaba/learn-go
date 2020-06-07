/**
@decs 结构体指针
@desc 结构体字段可以通过结构体指针来访问
@desc 指针间接的访问是透明的
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Object struct {
	X int
	Y int
}

func main() {
	i := Object{11, 33}
	m := make(map[string]interface{})
	fmt.Println(i) //{11 33}
	p := &i
	p.Y = 999
	fmt.Println(p)  //&{11 999}
	fmt.Println(*p) //指向的值

	j, _ := json.Marshal(p)
	_ = json.Unmarshal(j, &m)
	fmt.Println(m)
}
