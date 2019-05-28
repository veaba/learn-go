package _switch

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now() //2019-05-25 15:21:18.6070939 +0800 CST m=+0.003995301
	fmt.Println("xxxx", today)
	fmt.Println(today.Hour())

}
