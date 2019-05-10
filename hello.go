/**
@1 定义了包名，非注释的第一行，指明这个文件属于哪个包
@package main  一个可独立执行的程序。
@每个Go 都包含一个main包

*/
package main


/**
@1 fmt 告诉编译器，需要使用fmt包（函数，或者其他）
@2 fmt 包含格式化IO 的函数

*/
import "fmt"


/**
@1 开始执行函数，main，每一个可执行程序必须包含
@2 启动后 第一个执行的函数
@3 有init()函数执行，则会先init
*/
func main()  {
	test()
	fmt.Printf("hello,worl11d \n");
}
func init(){
	fmt.Printf("init,worl11d \n");
}

func test()  {
	fmt.Print("xxx\n");
}

// init,worlld
// xxx
// hello,worlld

//func main()  {
//	fmt.Printf("hello,world \n");
//}
//
