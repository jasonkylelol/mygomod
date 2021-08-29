package mygomod

import (
	"fmt"
	"time"
)

func printInOrder() {
	chA := make(chan string)
	chB := make(chan string)
	chC := make(chan string)
	for i := 0; i < 3; i++ {
		ch := chA
		nextCh := chB
		val := "A"
		switch i {
		case 1:
			ch = chB
			nextCh = chC
			val = "B"
		case 2:
			ch = chC
			nextCh = chA
			val = "C"
		}
		go func(goch, goNextCh chan string, goVal string) {
			for n := range goch {
				_ = n
				fmt.Println(goVal)
				time.Sleep(1 * time.Second)
				goNextCh <- "n"
			}
		}(ch, nextCh, val)
	}
	chA <- "n"
	for {
		time.Sleep(1 * time.Second)
	}
}
