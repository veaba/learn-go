/**
@desc 语言范围
*/
package main

import "fmt"

func main()  {
	nums := []int {998,65,6856,6}//定义一个切片
	sum:=0

	// 这个 "_"干嘛的？不需要索引，所以使用空白符 “_”省略
	for index,num:=range nums {
		fmt.Println("num:",num)
		fmt.Println("index:",index)
		sum+=num
	}

	fmt.Println("sum:",sum)
}
