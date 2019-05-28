/**
@desc 递归 recursion

*/
package recursion

import (
	"fmt"
)

func main() {
	var a = recursion(8)
	fmt.Println(a)
}

//
func recursion(num int) int {
	if num <= 1 {
		return 1
	}
	return recursion(num-1) + recursion(num-2)
}

// 1 1
// 2 2
// 3 3
