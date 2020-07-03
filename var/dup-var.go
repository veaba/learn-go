package main

import "fmt"

func main() {
	var a string
	var b1 = "b1"
	var b2 = "b2"
	var c = 2
	if c == 3 {
		a = b1
	} else {
		a = b2
	}
	fmt.Println(a)
}
