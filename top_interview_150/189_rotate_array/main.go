package main

import (
    "fmt"
    "reflect"
)

func rotate(nums []int, k int)  {
    
    n := len(nums)
    k = k % n

	buf := make([]int, k)

    for i, num := range nums[n-k:] {
		buf[i] = num
	}
    copy(nums[k:], nums[:n-k])
    copy(nums[:k], buf)
}


// Test cases
func runTests() {
    tests := []struct {
        input  []int
        k      int
        output []int
    }{
        {
            input:  []int{1, 2, 3, 4, 5, 6, 7},
            k:      3,
            output: []int{5, 6, 7, 1, 2, 3, 4},
        },
        {
            input:  []int{-1, -100, 3, 99},
            k:      2,
            output: []int{3, 99, -1, -100},
        },
        {
            input:  []int{1, 2, 3},
            k:      4, // k > len(input)
            output: []int{3, 1, 2},
        },
        {
            input:  []int{1},
            k:      10, // Single element array
            output: []int{1},
        },
    }

    for i, test := range tests {
        nums := make([]int, len(test.input)) // Create a copy to preserve original input
        copy(nums, test.input)
        rotate(nums, test.k)
        if !reflect.DeepEqual(nums, test.output) {
            fmt.Printf("Test %d FAILED: got %v, expected %v\n", i+1, nums, test.output)
        } else {
            fmt.Printf("Test %d PASSED\n", i+1)
        }
    }
}

func main() {
    runTests()
}
