/**
rand.Intn(100)

*/
package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("Lucky number:",rand.Intn(10))//1
	fmt.Println("Lucky number:",rand.Intn(100))//87 ?为什么一直是87
}
