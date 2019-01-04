
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
/*demo2*/
// const (
// 	i=1<<iota //左移 i=1<<0
// 	j=3<<iota //左移 j=3<<1
// 	t
// 	k=6<<iota
	
// )
// func main(){
// 	fmt.Println(i,j,t,k) //1
// }
/*demo3*/
// func main(){
// 	const a=3
// 	const b=10;
// 	var c uint=0
// 	c = a&b
// 	fmt.Println(c) //2
// 	fmt.Println(&c)
// 	c = a|b
// 	fmt.Println(c) //
// 	c =a ^ b
// 	fmt.Println(c)
// 	c= a<<2
// 	fmt.Println(c)
// 	c=a>>2
// 	fmt.Println(c)
// }
/*demo4 其他运算符*/
// func  main(){
// 	var a int =4
// 	var b int32
// 	var c float32
// 	var ptr *int
// 	/* 运算符实例 */
// 	fmt.Printf("%T",a)
// 	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
// 	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
// 	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
 
// 	/*  & 和 * 运算符实例 */
// 	ptr = &a    /* 'ptr' 包含了 'a' 变量的地址 */
// 	fmt.Printf("a 的值为  %d\n", a);
// 	fmt.Printf("*ptr 为 %d\n", *ptr);
// }
/*demo5*/
// func main(){
// 	numbers := [10]int{1,13,154,5,544,4,58,121}
// 	fmt.Println(numbers)
// 	for i,x:= range numbers {
// 		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
// 	 }   
// }
/*demo6*/
func main(){
	var a int =100;
	var b int =200;
	var c = max(a,b)
	fmt.Printf("%d\n",c)
}
// 必须对参数设置类型
func max(a1,b1 int) int {
	var res int
	if(a1>b1){
		res =a1
	}else{
		res =b1
	}
	return res
}
