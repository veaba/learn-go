/**
@io 包指定了io.Reader 接口，表示从数据流结尾读取
- 每次8字节的速度读取它的输出
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("你xxx挖到ad是好，傻等我打da54456d4傻!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v, err=%v, b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])

		if err == io.EOF {
			fmt.Println("exit 退出")
			break
		}
	}
}
