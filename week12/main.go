package main

import (
	"fmt"
)

func main() {

	s := []int{0, 0, 0, 0, 0}
	s[4] = 99
	s[2] = 11
	for _, value := range s {
		fmt.Println(value)
	}
	//fmt.Println("+-=")
	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	test := [3]string{"inha", "go", "student"}
	testS := test[:2]
	fmt.Println(len(testS))

	test2S := test[1:]
	fmt.Println(test2S[1])
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
