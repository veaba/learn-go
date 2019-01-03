
package main //定义包名，思考，类似C语言？为什么要这么做？

import "fmt"//告诉编译器，我要fmt包的函数，意思是类似 js 的 import {fmt} from main。包含格式化IO输入/输出函数
// var x,y int
// var (
// 	a int
// 	b bool
// )
// var c,d int =1,2
// var e,f =123,"hello"
// func main() { //程序开始执行的函数main 在这里是什么意思？可执行函数都必须包含这个 main！！！！？？？(o\~/o）
// 	_2,h:=123,"world"
// 	fmt.Println(_2,h)
// }
// func main(){
// 	const l  int =10;
// 	const w int =5
// 	var area int
// 	const a,b,c=1,false,"str"
// 	area = l * w
// 	fmt.Printf("面积为 : %d", area)//%d 此处何意？
// 	println()
// 	println(a,b,c)
// }
/**demo1*/
// func main(){
// 	const (
// 		a=iota    	//0 只是索引值，并不能认为是常量的int类型相加递增
// 		b			//1	
// 		c			//2
// 		d ="haha"	// haha
// 		e			//haha	=> "haha"+4
// 		f			//haha	=> "haha"+5
// 		g			//haha  => "haha"+6
// 		h=100		//?100
// 		i			// 为什么还是100
// 		j			// 为什么还是100
// 		k=iota		// 10
// 		l			// 11
// 	)
// 	fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l)
// }
// demo2
const (
	i=1<<iota //左移 i=1<<0
	j=3<<iota //左移 j=3<<1
	t
	k=6<<iota
	
)
func main(){
	fmt.Println(i,j,t,k) //1
}
/**
* @desc go
* // ->疑问
* 1、无法使用单引号
* // ->基本
	len(a)
* // -> fmt包
	.println(t)
	.Println(a,b)??
* // -> unsafe
	.Sizeof() ???
* // -> 基础结构
* 包声明
* 引入包
* 函数
* 变量
* 语句&表达式
* 注释
* // -> 标记
* 关键字
* 常量
* 字符串
* 符号
* fm
* .
* Println
* (
* "hello world"
* )
*  // -> 关键字
* break default func interface select case defer go map struct chan else goto package switch const fallthrougt if range type continue for imort return var 
*
* // -> 预定义标识符
* append bool byte cap close complex complex64 complex128 uint uint8 uint16 uint32 uint64 uintptr copy false float32 float64 imag int int8 int16 int32 iota len make new nil panic print println real recover string true 
*
* // -> 声明变量
* var age int; 
	- 必须空格区分
	- 被声明的变量，必须使用，对于局部作用域来说，但全局是允许声明可不使用。
	- 可在一行声明多个变量。var e,f =123,"hello" => g,h:=123,"world"，并行赋值|同时赋值
	- 不能对变量，重复声明
	- 交换两个变量的值，前提是类型相同 a,b=b,a
	- 空白标识符，_ 用于抛弃值 _,b=5,7. _ 
*
* // -> 值类型和引用类型
* 值类型-> (基本类型，存储在栈中，常量)  j=i 将i的值进行拷贝,&i 获取i的内存地址。每次的地址可能不一样
* 			int
			float
			bool
			string
* //=>常量
		常量用于枚举
		const {
			Unkown=0
			Female=1
			male=2
		}
		=>iota(特殊常量，可以被编译器修改的常量)???		
		---------------------------------
		|const (  	|		|const (	|
		|	a=iota	|		|	a=iota	|
		|	b=iota  |   =>	|	b		|
		|	c=iota	|		|	c		|
		|)			|		|)			|
		---------------------------------
		=>iota 用法
		package main
		import "fmt"
		func main(){
			const (
				a=iota
				b
				c
				d ="haha"
				e
				h
				i
				f=100
				g=iota
				i
			)
		}
* 引用类型->(复杂的数据结构，)r2=r1 只有引用被拷贝
*
* // ->数据类型
* 布尔类型-> 
			true 
			false
* 数字类型-> 
			整形int
				uint8 0-255
				uint16 0-65535
				uint32 0-4294967295
				uint64 0-188446744073709551615
				int8 -128-127
				int16 -32768-32767
				int32 -2147483648-2147483647
				int64 -9223372036854775808-9223372036854775807
			浮点float32、float64
				float32 IEEE-754 32
				float64 IEEE-754 64
				complex64 32位实数和虚数
				complex128 64位实数和虚数 (实数：有理数+无理数) （复数：实数+虚数）	
			其他 
				byte uint8
				rune ini32
				uint 32/64
				int 与uint
				uintptr 无符号整形，存放一个指针pointer
* 字符串类型-> 
			UTF-8编码标识的 
			unicode文本
* 派生类型-> 
*			指针类型 Pointer
			数组类型
			结构化类型 struct
			Channel 类型
			函数类型
			切片类型
			接口类型 interface
			Map 类型
*
*
*
* 操作符->
	:= -> (赋值操作符)
		var a =5;  => a:=5
		var b =false; => b:=false
*/