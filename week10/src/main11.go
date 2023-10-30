package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	greeting.Hello()

	fmt.Println(mymath.MyAbs(99))
	fmt.Println(mymath.MyAbs(-7))
	fmt.Println(mymath.MyPower(2, 10))

	greeting.Hi()
}
