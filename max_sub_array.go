package mygomod

import "math"

func maxSubArray(nums []int) int {
	val := 0
	max := nums[0]
	for _, num := range nums {
		val = int(math.Max(float64(val+num), float64(num)))
		max = int(math.Max(float64(val), float64(max)))
	}
	return max
}
