package mygomod

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	need, window := map[rune]int{}, map[rune]int{}
	for _, c := range ss1 {
		need[c]++
	}
	var left, right, valid int
	for right < len(ss2) {
		c1 := ss2[right]
		right++
		c1Val, ok := need[c1]
		if ok {
			window[c1]++
			if window[c1] == c1Val {
				valid++
			}
		}
		fmt.Printf("need:%v window:%v")
		for right-left >= len(ss1) {
			if valid == len(need) {
				return true
			}
			c2 := ss2[left]
			left++
			c2Val, ok := need[c2]
			if ok {
				window[c2]--
				if window[c2] == c2Val {
					valid--
				}
			}
		}
	}
	return false
}
