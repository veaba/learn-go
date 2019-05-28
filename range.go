/**
@desc 语言范围
*/
package main

import "fmt"

func main() {
	nums := []int{998, 65, 6856, 6} //定义一个切片
	sum := 0

	// 这个 "_"干嘛的？不需要索引，所以使用空白符 “_”省略 迭代
	for _, num := range nums {
		//fmt.Println("num:",num)
		//fmt.Println("index:",index)
		sum += num
	}

	fmt.Println("sum:", sum)

	//range 可以在map的键值对上。
	kvs := map[string]string{"who": "IG输了", "MSI2019": "舒服了"}

	fmt.Println(kvs) //map[MSI2019:舒服了 who:IG输了]
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}
}
