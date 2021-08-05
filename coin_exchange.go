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

func dpfuncNoRecurse(coins []int, amount int) int {
	dplist := []int{}
	for i := 0; i < amount+1; i++ {
		dplist = append(dplist, amount+1)
	}
	dplist[0] = 0
	for idx := range dplist {
		for _, coin := range coins {
			if idx-coin < 0 {
				continue
			}
			if dplist[idx] > 1+dplist[idx-coin] {
				dplist[idx] = 1 + dplist[idx-coin]
			}
		}
		fmt.Println(dplist)
	}
	if dplist[amount] == amount+1 {
		return -1
	}
	return dplist[amount]
}

func coinChange(coins []int, amount int) int {
	// dpMap := map[int]int{}
	// return dpfunc(coins, amount, &dpMap)
	return dpfuncNoRecurse(coins, amount)
}
