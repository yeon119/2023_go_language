package main

import "fmt"

func main() {

	s := []int{0, 0, 0, 0, 0}
	for _, value := range s {
		fmt.Println(value)
	}

	// s := make([]int, 5)
	// for _, value := range s {
	// 	fmt.Println(value)
	// }

	// var s []int
	// s = make([]int, 5)
	// for _, value := range s {
	// 	fmt.Println(value)
	// }
}
