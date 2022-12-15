package main

import (
	"fmt"
	"sync"
	"time"
)

func rpcCallV2(b, p int) int {
	time.Sleep(1 * time.Second)
	return b*10 + p
}

func concurrencyV2() {
	fmt.Println("Hello World")

	// var mutex sync.Mutex
	var wg sync.WaitGroup
	reslist := []int{}
	resmap := map[int]int{}
	//var results []string

	for board := 0; board < 10; board++ {
		for proj := 0; proj < 10; proj++ {

			wg.Add(1)
			//因为是闭包函数，其实现在保证了函数内引用变量的生命周期与函数的活动时间相同
			go func(b, p int) {
				defer wg.Done()

				val := rpcCallV2(b, p)

				reslist = append(reslist, val)
				resmap[val] = val
			}(board, proj)
		}
	}

	wg.Wait()

	fmt.Println("reslist len:", len(reslist))
	fmt.Println("resmap len:", len(resmap))
}
