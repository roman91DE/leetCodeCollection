package main

import (
	"fmt"
)

// canConstruct checks if ransomNote can be constructed from magazine
func canConstruct(ransomNote string, magazine string) bool {

	// create the counter object

	// optimize 4 faster runtime
	// letters := make(map[rune]int, len(magazine))

	// optimize 4 memory efficiency
	letters := make(map[rune]int)

	for _, c := range magazine {
		letters[c]++
	}

	// check if ransomNote can be constructed from letters
	for _, cc := range ransomNote {
		if v, exists := letters[cc]; exists && v > 0 {
			letters[cc]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	testCanConstruct()
}

// Tests for canConstruct
func testCanConstruct() {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"hello", "helloworld", true},
		{"abc", "abcd", true},
		{"abc", "ab", false},
		{"", "abc", true},
		{"abc", "", false},
	}

	for i, test := range tests {
		result := canConstruct(test.ransomNote, test.magazine)
		if result != test.expected {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, test.expected, result)
		} else {
			fmt.Printf("Test case %d passed\n", i+1)
		}
	}
}
