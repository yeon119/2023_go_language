package main

import "fmt"

func main() {

	s := make([]int, 5)
	for _, value := range s {
		fmt.Println(value)
	}

	// var s []int
	// s = make([]int, 5)
	// for _, value := range s {
	// 	fmt.Println(value)
	// }
}
