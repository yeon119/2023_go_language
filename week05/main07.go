package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("점수를 입력하세요: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //option2
	if err != nil {                       //conditionals
		log.Fatal(err)
	}
	inputScore = string.TrimSpace(inputScore)
	score, err := strconv.ParseFloat
	fmt.Println(input)
}
