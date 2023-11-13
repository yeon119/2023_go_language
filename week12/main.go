package main

import (
	"fmt"
)

func main() {
	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	c := append(a, "x", "y")
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	as[1] = "z"
}
