/**
@desc 结构体
@TODO 数组可以存储同一个类型的数组
@TODO 数组如何支持结构体 arr=[999,"899",false]
@TODO 那只是属性而已，怎么去调用方法？？
@struct 结构体可以定义不同的数据类型
@ 一系列具有同类型或者不同类型的数据构成的数据集合
@类似JavaScript中的Object 对象
*/
package main

import "fmt"

type Books struct {
	name   string
	author string
	page   int
}

func main() {
	fmt.Println(Books{"《西游记》", "吴承恩", 9999})

	fmt.Println(Books{name: "《红楼梦》", author: "曹雪芹", page: 544})

	fmt.Println(Books{name: "《三国演义》", page: 666})

	var book1 Books
	book1.name = "《水浒传》"

	fmt.Println(book1) //{《水浒传》  0} 为什么 0 也要显示呢？0 或者空

	fmt.Println(book1.name)
	fmt.Println(book1.author)

	// 结构体指针
	structPtr()
}

func structPtr() {
	var book2 Books
	book2.author = "小李子"

	printBook(&book2)

	pointBook(book2)
}

func pointBook(books Books) {
	fmt.Println(books) //{ 小李子 0}
}

func printBook(books *Books) {
	fmt.Println(books.author)
	fmt.Println(books) //&{ 小李子 0}
}
