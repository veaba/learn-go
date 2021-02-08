package main

import (
	"fmt"
	"strconv"
)

func main3() {
	var i = "8"
	fmt.Println(strconv.ParseInt(i, 10, 32))
}
