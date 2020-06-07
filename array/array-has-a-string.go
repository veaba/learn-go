package main

import (
	"fmt"
	"strings"
)

func main() {

	array1 := [11]string{"数学", "语文", "英语", "化学", "生物", "物理", "历史", "其他", "语文思维训练", "数理逻辑训练", "学习力训练"}
	fmt.Println("json-pr===>", array1)

	str := "语文阅读理解80几分"

	for k, v := range array1 {
		fmt.Println("==>", k, v)
	}

	x := strings.Contains(str, "语文")
	fmt.Println("====>", x)

	fmt.Println(strings.Contains("数文", "数学"))

}
