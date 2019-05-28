/**

@desc //fmt.Printf("%i",Big)//怎
*/
package numeric_const

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	fmt.Println("xxx:", x) //2
	return x*10 + 1
}

func needFloat(x float64) float64 {
	fmt.Println("ooo:", x) //1.2676506002282294e+30
	return x * 0.1
}

func main() {

	//fmt.Println(Small)
	fmt.Println(needInt(Small))   //21
	fmt.Println(needFloat(Small)) //0.2
	fmt.Println(needFloat(Big))   //1.2676506002282295e+29
}
