package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")
	word := words[len(words)-1]
	return len(word)
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
