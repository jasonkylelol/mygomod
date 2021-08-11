package mygomod

import (
	"fmt"
	"reflect"
)

func printMap(rmap map[rune]int) string {
	sval := ""
	for key, val := range rmap {
		sval += fmt.Sprintf("%v:%v ", string(key), val)
	}
	return sval
}

func checkInclusion(s1 string, s2 string) bool {
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	need, window := map[rune]int{}, map[rune]int{}
	for _, c := range ss1 {
		need[c]++
	}
	var left, right int
	for right < len(ss2) {
		c1 := ss2[right]
		right++
		_, ok := need[c1]
		if ok {
			window[c1]++
		}
		// fmt.Printf("need:%v window:%v c1:%v\n", printMap(need), printMap(window), string(c1))
		for right-left >= len(ss1) {
			if reflect.DeepEqual(window, need) {
				return true
			}
			c2 := ss2[left]
			left++
			_, ok := need[c2]
			if ok {
				window[c2]--
			}
		}
	}
	return false
}
