package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dice := rand.Intn(6) + 1 // 0~5에 1을 더한것
	fmt.Println(dice)
}
