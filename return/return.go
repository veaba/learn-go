/**

- 可以返回任意数量的返回值——多返回值
- JavaScript语言中则不能这样，需要放在一个对象里面

*/
package _return

import "fmt"

func Return(x, y string) (string, string) {
	return x, y
}

func main() {
	hello, world := Return("hello", "world")
	fmt.Println(hello, world)
}
