package main

import (
	"fmt"
)

func shift(s []rune) []rune {
	return append(s[1:], s[0])
}


func rotateString(s string, goal string) bool {
    if s == goal {
		return true
	}
	if len(s) != len(goal) {
		return false
	}
	shifted := []rune(s)

	for i := 0; i < len(s); i++ {
		shifted = shift(shifted)
		if string(shifted) == goal {
			return true
		}
	}
	return false
}

func main(){
	s := "abcd"
	goal := "cdab"

	fmt.Printf("Check True: %t\n", rotateString(s,goal))

	s = "bcd"
	goal = "cab"

	fmt.Printf("Check False: %t\n", rotateString(s,goal))
}