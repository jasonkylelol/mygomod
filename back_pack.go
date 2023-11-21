package mygomod

import "math"

/*

    N = 3, W = 4
    wt = [2, 1, 3]
    val = [4, 2, 3]
*/

func backPack(cap, num int, wt, val []int) int {

	res := [][]int{}
	for i := 0; i < num+1; i++ {
		subcap := []int{}
		for j := 0; j < cap+1; j++ {
			subcap = append(subcap, 0)
		}
		res = append(res, subcap)
	}
	for i := 1; i <= num; i++ {
		for w := 1; w <= cap; w++ {
			if w-wt[i-1] < 0 {
				// 
				res[i][w] = res[i-1][w]
			} else {
				// 
				res[i][w] = int(math.Max(
					float64(res[i-1][w-wt[i-1]]+val[i-1]),
					float64(res[i-1][w])))
			}
		}
	}
	return res[num][cap]
}
