package mygomod

import "sort"

func innerTwoSum(nums []int, start, target int) [][]int {
	low := start
	high := len(nums) - 1
	res := [][]int{}
	for low < high {
		lval, hval := nums[low], nums[high]
		sum := lval + hval
		if sum < target {
			low++
		} else if sum > target {
			high--
		} else {
			res = append(res, []int{nums[low], nums[high]})
			for low < high {
				if lval != nums[low] {
					break
				}
				low++
			}
			for low < high {
				if hval != nums[high] {
					break
				}
				high--
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for idx := 0; idx < len(nums); idx++ {
		val := nums[idx]
		target := 0 - val
		twosumlist := innerTwoSum(nums, idx+1, target)
		for _, twosums := range twosumlist {
			twosums = append(twosums, val)
			res = append(res, twosums)
		}
		for idx < len(nums)-1 {
			if nums[idx] != nums[idx+1] {
				break
			}
			idx++
		}
	}
	return res
}
