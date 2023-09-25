package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("점수를 입력하세요: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') //option2
	if err != nil {                       //conditionals
		log.Fatal(err)
	}

	fmt.Println(input)
}
