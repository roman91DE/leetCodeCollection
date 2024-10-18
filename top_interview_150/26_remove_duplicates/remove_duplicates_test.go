package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
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
			name:      "Case 2",
			input:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:  []int{0, 1, 2, 3, 4},
			expectedK: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of the input slice
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			// Run removeDuplicates
			k := removeDuplicates(inputCopy)

			// Check the result
			if k != tc.expectedK {
				t.Errorf("expected k: %d, got: %d", tc.expectedK, k)
			}

			for i := 0; i < k; i++ {
				if inputCopy[i] != tc.expected[i] {
					t.Errorf("expected element at index %d: %d, got: %d", i, tc.expected[i], inputCopy[i])
				}
			}
		})
	}
}
