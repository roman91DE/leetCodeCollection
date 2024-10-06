package main

import (
	"fmt"
	"sort"
	"reflect"
)

func merge1(nums1 []int, m int, nums2 []int, n int) {
	// Copy the elements from nums2 into nums1 starting from index m
	copy(nums1[m:], nums2[:n])
	// Sort the entire nums1 slice in place
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := 0, 0

	for i1 < m+n && i2 < n {
		if i1 >= m || nums2[i2] < nums1[i1] {
			copy(nums1[i1+1:], nums1[i1:])
			nums1[i1] = nums2[i2]
			i2++
			m++ // Increase m to reflect the addition of the new element
		} else {
			i1++
		}
	}
	
	// Copy any remaining elements from nums2 to nums1
	if i2 < n {
		copy(nums1[m:], nums2[i2:])
	}
}



// testMerge runs test cases to verify the merge function works as expected.
func testMerge() {
	tests := []struct {
		nums1   []int
		m       int
		nums2   []int
		n       int
		expected []int
	}{
		{
			nums1:   []int{1, 2, 3, 0, 0, 0},
			m:       3,
			nums2:   []int{2, 5, 6},
			n:       3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:   []int{4, 0, 0, 0},
			m:       1,
			nums2:   []int{1, 2, 3},
			n:       3,
			expected: []int{1, 2, 3, 4},
		},
		{
			nums1:   []int{0},
			m:       0,
			nums2:   []int{1},
			n:       1,
			expected: []int{1},
		},
	}

	for i, test := range tests {
		// Make a copy of nums1 to avoid modifying the original test case data
		nums1Copy := make([]int, len(test.nums1))
		copy(nums1Copy, test.nums1)

		// Run the merge function
		merge2(nums1Copy, test.m, test.nums2, test.n)

		// Check if the result matches the expected value
		if !reflect.DeepEqual(nums1Copy, test.expected) {
			fmt.Printf("Test case %d failed: got %v, expected %v\n", i+1, nums1Copy, test.expected)
		} else {
			fmt.Printf("Test case %d passed\n", i+1)
		}
	}
}

func main() {
	// Run the test cases
	testMerge()
}