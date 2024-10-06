package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	idx := sort.SearchInts(nums, val)
	swapIdx := len(nums) - 1

	for ; idx < swapIdx; idx = sort.SearchInts(nums[:swapIdx], val) {
		for ; nums[swapIdx] == val; swapIdx-- {
			break
		}
		nums[idx], nums[swapIdx] = nums[swapIdx], nums[idx]
		swapIdx--
	}

	return idx
}
func main() {
	// Example 1
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	k1 := removeElement(nums1, val1)
	if k1 == 2 && equal(nums1[:k1], []int{2, 2}) {
		fmt.Println("Test case 1 passed")
	} else {
		fmt.Println("Test case 1 failed")
	}

	// Example 2
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	k2 := removeElement(nums2, val2)
	if k2 == 5 && equal(nums2[:k2], []int{0, 0, 1, 3, 4}) {
		fmt.Println("Test case 2 passed")
	} else {
		fmt.Println("Test case 2 failed")
	}

	// Additional test case 3 (edge case: empty array)
	nums3 := []int{}
	val3 := 1
	k3 := removeElement(nums3, val3)
	if k3 == 0 && equal(nums3[:k3], []int{}) {
		fmt.Println("Test case 3 passed")
	} else {
		fmt.Println("Test case 3 failed")
	}

	// Additional test case 4 (edge case: no elements to remove)
	nums4 := []int{1, 2, 3, 4, 5}
	val4 := 6
	k4 := removeElement(nums4, val4)
	if k4 == 5 && equal(nums4[:k4], []int{1, 2, 3, 4, 5}) {
		fmt.Println("Test case 4 passed")
	} else {
		fmt.Println("Test case 4 failed")
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
