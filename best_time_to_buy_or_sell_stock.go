package mygomod

import "math"

/*

input: [7,1,5,3,6,4]
output: 5
buy day 2 (price = 1), sell day 5(price = 6), max profit = 6-1 = 5
*/

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profit
}
