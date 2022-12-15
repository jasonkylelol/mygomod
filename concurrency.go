package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrency() {
	fmt.Println("Hello World")
	boards := []string{"board1", "board2"}
	projects := []string{"aaaaa", "bbbbb", "ccccc"}

	var wg sync.WaitGroup
	var results []string
	resmap := map[string][]string{}
	n := time.Now()
	_ = boards
	_ = projects
	_ = n
	_ = resmap

	for board := 0; board < 10; board++ {
		// var sss string
		for proj := 0; proj < 10; proj++ {
			wg.Add(1)
			go func(b, p int) {
				defer wg.Done()

				// fmt.Println("board:", b, "project:", p)
				//time.Sleep(1 * time.Second)
				sss := fmt.Sprintf("%v_%v", b, p)
				results = append(results, sss)
				// resmap[b] = append(resmap[b], p)
				// time.Sleep(1 * time.Second)
			}(board, proj)
			//wg.Wait
			// results = append(results, sss)
		}
	}
	wg.Wait()

	fmt.Println("results len:", len(results))
	// fmt.Println("resmap:", resmap)
}
