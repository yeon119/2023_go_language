package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer_word := strings.NewReplacer("#", "o")
	fixedWord := replacer_word.Replace(broken)
	fmt.Println(fixedWord)
}
