package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("나는 1부터 100 까지의 숫자중 하나를 고를 것 입니다")
	fmt.Println("추측하실수 있나요?")
	fmt.Println(target)
	
	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++
	fmt.Println("너는 " ,10-guesses, "만큼의 기회가 남았따" )

	fmt.Println("추측해봐:")
	inpur, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)

	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if guess < target {
		fmt.Println("너듸 답은 작다")
	} else if guess >target {
		fmt..Println("저런 너의 답은 높다")
	}
	

}
