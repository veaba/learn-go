package type_conversion

import (
	"fmt"
)

func main() {
	x, y := 3, 4
	//f := math.Sqrt(float64((x*x+y*y)))
	z := int(x*x + y*y)
	fmt.Println(x, y, z)

	i := 42
	f64 := float64(i)
	fmt.Println(i, f64)
}
