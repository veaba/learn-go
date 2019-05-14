package main

import "fmt"

const (
	a=iota
	b=iota
	c=iota
	d=8999
	e
	f
	g=100
	i
	j
	l=iota
	m
	n
)

// 左移
const (
	a1=1<<iota	//1<<0
	a2=2<<iota	//2<<1
	a3			//2<<2
	a4			//2<<3
)

const (
	b1=1<<iota 	//1<<0
	b2=3<<iota 	//3<<1
	b3         	//3<<2
	b4			//3<<3
	b5=iota		//4
)
func foo()  {
	fmt.Println(a,b,c)
	fmt.Println(d,e,f,g)//d 被改变了，后续在重新iota之前都等于这个值
	fmt.Println(i,j) //g 被改变了，后续在l=iota之前都等于100
	fmt.Println(l,m,n) //l=iota 被恢复，开始恢复次序
}

func leftMove1()  {
	fmt.Println(a1,a2,a3,a4)//1 4 8 16
}
func leftMove2()  {
	fmt.Println(b1,b2,b3,b4,b5)//1 6 12 24
}

func main()  {
	//foo()
	leftMove1()
	leftMove2()
}
