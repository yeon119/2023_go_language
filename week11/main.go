package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	primes := [3]int{2, 3, 5} // 초기화를 배열 리터럴로

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}
	test := [5]bool{true, true, true}
	fmt.Println(test[3])
	fmt.Println((test))

	// i := 0
	// for i < 6 { // run time error !!
	// 	fmt.Println(test[i])
	// 	i++
	// }
	i := 0
	for i < len(test) { // len 함수를 이용하여 panic을/를 방지
		fmt.Println(test[i])
		i++
	}

	for prime := range primes { // index만 출력
		fmt.Println(prime)

	}

	// 	for idx ,prime := range primes { // idx를 출력하지않아 컴파일 에러
	// 		fmt.Println(prime)
	// }

	for _, prime := range primes { // index만 출력
		fmt.Println(prime)
	}
}
