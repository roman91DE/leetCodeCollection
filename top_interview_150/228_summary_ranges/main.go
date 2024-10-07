package main

import (
	"fmt"
)

type Interval struct {
	start int
	end   int
}

func (i Interval) toString(nums []int) string{
	if i.start == i.end {
		return fmt.Sprintf("%v", nums[i.start])
	} else {
		return fmt.Sprintf("%v->%v", nums[i.start], nums[i.end])
	}
}


func summaryRanges(nums []int) []string {

	if len(nums) == 0{
		return []string{}
	}
	
	intervals := make ([]Interval, 0)
	startBuffer := 0

	for i,v := range nums {

		if i == len(nums) -1 {
			// reached the end of the slice
			break
		} else if (v+1 == nums[i+1]) {
			// intervall keeps running
			continue
		} else {
			// reached the end of the intervall
			intervals = append(intervals, Interval{startBuffer, i})
			startBuffer = i+1
		}
	}
	intervals = append(intervals, Interval{startBuffer, len(nums)-1})

	intervalStrings := make([]string, len(intervals))
    for idx, interval := range intervals {
        intervalStrings[idx] = interval.toString(nums)
    }

	return intervalStrings
}



func main() {
	// Test cases from the challenge description
	testCases := []struct {
		input    []int
		expected []string
	}{
		{input: []int{0, 1, 2, 4, 5, 7}, expected: []string{"0->2", "4->5", "7"}},
		{input: []int{0, 2, 3, 4, 6, 8, 9}, expected: []string{"0", "2->4", "6", "8->9"}},
		{input: []int{}, expected: []string{}}, // Empty input case
	}

	for _, testCase := range testCases {
		output := summaryRanges(testCase.input)
		fmt.Printf("Input: %v\nOutput: %v\nExpected: %v\n\n", testCase.input, output, testCase.expected)
	}
}
