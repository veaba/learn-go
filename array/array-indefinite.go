/**
* 不定长的数组
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var indefinite_array []string
	for i := 0; i < 9; i++ {
		fmt.Println("i==>", i)
		indefinite_array = append(indefinite_array, strconv.Itoa(i)+"→")
	}
	fmt.Println("indefinite_array===>", indefinite_array)

}


