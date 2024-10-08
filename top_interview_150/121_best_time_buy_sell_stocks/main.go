package main

import (
	"fmt"
)

// maxProfit function calculates the maximum profit that can be achieved
func maxProfit(prices []int) int {

	profit, temp := 0, 0

	for buyIdx, buyPrice := range prices {
		for _, sellPrice := range prices[buyIdx:] {
			if sellPrice > buyPrice {
				temp = sellPrice - buyPrice
				if temp > profit {
					profit = temp
				}
			}

		}
	}
	return profit
}

func main() {
	// Test cases
	testCases := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4}, // Increasing prices
		{[]int{5, 4, 3, 2, 1}, 0}, // Decreasing prices
		{[]int{3, 3, 3, 3, 3}, 0}, // All prices are the same
		{[]int{}, 0},              // Empty array
	}

	for i, testCase := range testCases {
		result := maxProfit(testCase.prices)
		if result == testCase.expected {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: Input: %v, Expected: %d, Got: %d\n", i+1, testCase.prices, testCase.expected, result)
		}
	}
}
