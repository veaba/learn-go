/**
@desc go 语言中的 指针
@空指针，一个指针被定义后，没有分配到任何变量。值为nil
@&x  取变量的指针地址
@*x 取指针指向的内容
@指针的指针 指针变量存放的又是另外一个指针变量的地址，称 这个指针变量 位指向指针指针变量
nil => null None nil NULL 都指代零值或空值
*/
package main

import "fmt"

func viodPtr() {
	var prt *int

	if prt == nil {
		fmt.Println("空指针")
	}

	if prt != nil {
		fmt.Println("不等于空指针")
	}
	fmt.Println(prt)
}

func prtArray() {
	const MAX = 3
	a := []int{10, 30, 90}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] //整数地址复制给指针数组
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
	fmt.Println(a)
	fmt.Println(ptr) //思考既然知道指针，如何知道这个指针上面的值是多少？加一个* 号

	fmt.Println(*ptr[0])
	// 保存数组
}

func prtPrt() {
	var a int //0
	var ptr *int
	var pptr **int

	a = 3333
	ptr = &a //a的指针地址给ptr这个变量

	pptr = &ptr //ptr 指针的又赋值给pptr

	fmt.Printf("变量a = %d\n", a)                  //3333
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)         //3333
	fmt.Printf("指向指针的指针变量 **pptr= %d\n", **pptr) //3333
	fmt.Printf("指向指针的指针变量 *pptr = %d\n", *pptr)  //824634122496 What??
}
func ptrFunc() {
	var a = 666
	var b = 999

	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var pointer = 8888
	fmt.Println(pointer)  //888
	fmt.Println(&pointer) //0xc00000a0a8

	//空指针
	viodPtr() // <nil>

	//指针数组
	prtArray()

	//指向 指针的指针
	prtPrt()

	// 指针作为函数参数
	ptrFunc()
}
