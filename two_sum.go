package mygomod

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
