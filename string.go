package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "xxooo"
	f := "oo"
	t := "xx"
	b := strings.HasPrefix(s,f)//prefix 前缀
	fmt.Println(b,"\n")//false
	b = strings.HasPrefix(s,t)//prefix 前缀
	//fmt.Println(b)//true

	b = strings.HasSuffix(s,f) //stuffix 后缀
	//fmt.Println(b)//true

	b = strings.HasSuffix(s,t) //stuffix 后缀
	//fmt.Println(b)//false

	//test()
	//test1()
	//test3()
	test4()
}

func test4() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)//[The quick brown fox jumps over the lazy dog]
	fmt.Println(sl)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		//fmt.Printf("%s - 0.0 ", val)
		//fmt.Println(val)
		fmt.Printf("%s ",val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")

	fmt.Println(sl2)
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

func test3() {
	var origS string = "Hi there! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}

func test1() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}

func test() {
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}




