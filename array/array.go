package main

import "fmt"

func arr() {
	fmt.Println("I am arr")
}

func main() {
	arr()
	var arr = [9]float32{9.09, 9.6}

	fmt.Println(arr) //[9.09 9.6 0 0 0 0 0 0 0]

	var arr1 = [...]float32{5454, 545, 4545, 54}

	fmt.Println(arr1)      //5454,545,4545,54
	fmt.Println(len(arr1)) //4
	fmt.Println(len(arr))  //9

	arr1[3] = 99999

	fmt.Println(arr1) //[5454 545 4545 99999]

	// 不设定类型
	//var arr2 = [...]float32,float32,string{54545,54,"xx"}
	//
	//fmt.Println(len(arr3))
}
