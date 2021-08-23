package mygomod

import "math"

/*
给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少？
    举个简单的例子，输入如下：
    N = 3, W = 4
    wt = [2, 1, 3]
    val = [4, 2, 3]
    算法返回 6，选择前两件物品装进背包，总重量 3 小于 W，可以获得最大价值 6。
*/

func backPack(cap, num int, wt, val []int) int {
	// base case 已初始化
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
				// 这种情况下只能选择不装入背包
				res[i][w] = res[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				res[i][w] = int(math.Max(
					float64(res[i-1][w-wt[i-1]]+val[i-1]),
					float64(res[i-1][w])))
			}
		}
	}
	return res[num][cap]
}
