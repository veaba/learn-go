package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)      //windows
	fmt.Println(runtime.GOARCH)    //amd64
	fmt.Println(runtime.Version()) //go1.12.1
}
