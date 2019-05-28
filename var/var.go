package _var

import "fmt"

func main() {

	// 变量的使用
	//var a= "xxoo"
	//var b string = "hello "
	//fmt.Println(a)
	//fmt.Println(b + "world")
	//
	//fmt.Println(a, b, "台湾~")

	// 没有初始化值的变量
	/*var c int
	fmt.Println(c)

	var f  bool
	fmt.Println(f)

	if f{
		fmt.Println(f,!f,!!f)
	}else {
		fmt.Println(f,!f,!!f)
	}

	if c==0 {
		fmt.Println(c)
	}*/

	// 为nil 的声明

	/*var a1 int
	var a2 float64
	var a3 bool
	var a4 string
	fmt.Println(a1, a2, a3, a4)
	//fmt.Printf("%v %v %v %q \n",a1, a2, a3, a4)
	fmt.Printf("%v %v %v %q \n",a1, a2, a3, a4)*/

	//错误 :=左侧是没有被声明过的变量
	/*var val int
	val :=2
	fmt.Println(val)*/

	//第三种 省略var 关键字变量声明

	var val1 = "a"

	var val2 = "b"

	val1 = val2
	val2 = val1

	//var val3 = val2
	//
	//var val4 = "99999"
	//var val5 = val4

	fmt.Println(val1, val2)
	fmt.Println("&val1", &val1) //
	fmt.Println("&val2", &val2) //内存地址可能一样的
	//fmt.Println("&val3",&val3) //内存地址和val2不一样
	//fmt.Println("&val4",&val4) //内存地址可能一样的
	//fmt.Println("&val5",&val5) //内存地址可能一样的

	// 多变量声明
	/*var (
		name  = 222
		types = 22
	)

	fmt.Println(name)
	fmt.Println(types)*/
}
