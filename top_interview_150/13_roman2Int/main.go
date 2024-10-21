package main

import "fmt"

// Recursive function to convert Roman numeral string to integer
func recsum(s string, acc int, romanMap map[rune]int) int {
	// Base case: if the string is empty
	if len(s) == 0 {
		return acc
	} else if len(s) == 1 {
		// If only one character remains, add its value to the accumulator
		return acc + romanMap[rune(s[0])]
	} else {
		// Get the first and second characters
		first, second := romanMap[rune(s[0])], romanMap[rune(s[1])]
		if first < second {
			// If the first character is less than the second, subtract and skip the next two characters
			acc += second - first
			return recsum(s[2:], acc, romanMap)
		} else {
			// Otherwise, add the first character and continue with the next
			acc += first
			return recsum(s[1:], acc, romanMap)
		}
	}
}

// Main function to convert Roman numeral string to integer
func RomanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Start the recursion with an accumulator of 0
	return recsum(s, 0, romanMap)
}

func main() {
	// Example usage
	fmt.Println(RomanToInt("IX"))       // Should print 9
	fmt.Println(RomanToInt("MCMXCIV"))  // Should print 1994
}
