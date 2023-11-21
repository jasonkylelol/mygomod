package mygomod

/*
two interger array: nums1 和 nums2, m & n represents the number of items

merge nums2 to nums1, ascending order

input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
output: [1,2,2,3,5,6]

*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1
	for p1 >= 0 || p2 >= 0 {
		var cur int
		if p1 < 0 {
			cur = nums2[p2]
			p2--
		} else if p2 < 0 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] < nums2[p2] {
			cur = nums2[p2]
			p2--
		} else {
			cur = nums1[p1]
			p1--
		}
		nums1[p3] = cur
		p3--
	}
}
