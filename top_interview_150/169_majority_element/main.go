package main

import (
	"fmt"
)

// majorityElement returns the majority element in an array.
func majorityElement(nums []int) int {
	counter := make(map[int]int, len(nums))

	for _, num := range nums {
		counter[num]++
	}

	item, count := nums[0], 1

	for k, v := range counter {
		if v > count {
			item, count = k, v
		}
	}
	return item

}


func main() {
	// Test cases
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{3, 2, 3}, expected: 3},
		{input: []int{2, 2, 1, 1, 1, 2, 2}, expected: 2},
	}

	for i, testCase := range testCases {
		result := majorityElement(testCase.input)
		if result == testCase.expected {
			fmt.Printf("Test case %d passed.\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %d, got %d\n", i+1, testCase.expected, result)
		}
	}
}
