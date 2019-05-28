package math_Nextafter

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("%g xx", math.Nextafter(2, 3)) //2.0000000000000004
	fmt.Println("%g xx", math.Nextafter(2, 1)) //1.9999999999999998
	fmt.Println("%g xx", math.Nextafter(2, 2)) //2
}
