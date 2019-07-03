/**
@desc 命令行参数
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("A:", args)
	fmt.Println("0:", args[0])
	fmt.Println("1:", args[1])
}
