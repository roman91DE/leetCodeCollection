package main

import "fmt"

// very slow version that copies the whole sub array after each duplication found
func removeDuplicatesSlow(nums []int) int {
	i, j := 0, len(nums)
	for i < j-1 {
		if nums[i] == nums[i+1] {
			copy(nums[i:], nums[i+1:])
			j--
			continue
		}
		i++
	}
	return j
}

// removes duplicates from input slice and returns the index of the last unique item
func removeDuplicates(nums []int) int {

	if len(nums) < 1 { // this isnt necessary for the leetcode challenge
		return 0
	}
	k := 1 // insertion point for new unique items

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}

	}
	return k
}

func main() {
	// Test cases
	testCases := []struct {
		name      string
		input     []int
		expected  []int
		expectedK int
	}{
		{
			name:      "Case 1",
			input:     []int{1, 1, 2},
			expected:  []int{1, 2},
			expectedK: 2,
		},
		{
			name:  "Case 2",
			input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},

			expected:  []int{0, 1, 2, 3, 4},
			expectedK: 5,
		},
	}

	for _, tc := range testCases {
		// Make a copy of the input slice
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		// Run removeDuplicates
		k := removeDuplicates(inputCopy)

		// Print results
		fmt.Printf("Test Case: %s\n", tc.name)
		fmt.Printf("Expected k: %d, Got k: %d\n", tc.expectedK, k)
		fmt.Printf("Modified Array: %v\n", inputCopy[:k])
		fmt.Printf("Expected Array: %v\n\n", tc.expected)
	}
}
