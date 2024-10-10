package main

import (
	"fmt"
)

// my first implementation brute forces all combinations and is
// therefore too slow on large test cases

// func maxProfit(prices []int) int {
// 	profit, temp := 0, 0
// 	for buyIdx, buyPrice := range prices {
// 		for _, sellPrice := range prices[buyIdx:] {
// 			if sellPrice > buyPrice {
// 				temp = sellPrice - buyPrice
// 				if temp > profit {
// 					profit = temp
// 				}
// 			}
// 		}
// 	}
// 	return profit
// }

// maxProfit function calculates the maximum profit that can be achieved
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	// idea: transform array to price changes relative to the previous price
	priceDiffs := make([]int, len(prices))
	priceDiffs[0] = 0

	for i, val := range prices[1:] {
		priceDiffs[i+1] = val - prices[i]
	}

	// kardane's algorithm finds the largest sum over all contigous sub-arrays
	return kardaneAlgorithm(priceDiffs)
}

// kardane's algorithm finds the largest sum of a contigous subarray
func kardaneAlgorithm(diffs []int) int {
	// 0 also reflects negative returns in the challenge
	bestSum, currentSum := 0, 0

	for _, val := range diffs {
		// current sum is either the current value or the previous rolling sum + the current value
		currentSum = max(val, currentSum+val)
		// save best results
		bestSum = max(bestSum, currentSum)
	}
	return bestSum
}

func main() {
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
		{[]int{2, 4, 1}, 2},
		{[]int{2, 11, 1, 4, 7}, 9},
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
