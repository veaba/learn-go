package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := "hello world"
	fmt.Println(typeof(v))
}
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
