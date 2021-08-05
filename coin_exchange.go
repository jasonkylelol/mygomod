package mygomod

import "fmt"

func dpfunc(coins []int, amount int, dpMap *map[int]int) int {
	val, ok := (*dpMap)[amount]
	if ok {
		return val
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := amount + 1
	fmt.Printf("amount:%v res:%v\n", amount, res)
	for _, coin := range coins {
		sub := dpfunc(coins, amount-coin, dpMap)
		if sub == -1 {
			continue
		}
		if res > 1+sub {
			res = 1 + sub
		}
	}
	val = res
	if res == amount+1 {
		val = -1
	}
	(*dpMap)[amount] = val
	return val
}

func coinChange(coins []int, amount int) int {
	dpMap := map[int]int{}
	return dpfunc(coins, amount, &dpMap)
}
