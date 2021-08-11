package mygomod

/*
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	low, high := 0, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum < target {
			low++
		} else if sum > target {
			high--
		} else {
			return []int{low, high}
		}
	}
	return []int{}
}
*/

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	rets := []int{}
	var right int
	for right < len(nums) {
		rval := nums[right]
		var left int
		for left < right {
			lval := nums[left]
			if lval+rval == target {
				rets = append(rets, left, right)
				return rets
			}
			left++
		}
		right++
	}
	return []int{}
}
