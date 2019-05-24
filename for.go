package main

import "fmt"

func main()  {
	sum :=1
	for sum <64{
		fmt.Println(sum)
		sum+=sum
	}
	fmt.Println("end:",sum)
}
