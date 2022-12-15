package main

import (
	"fmt"
)

func myDefer() {
	id := 1
	defer func() {
		fmt.Println("defer:", id)
	}()
	id = 2
	fmt.Println("id:", id)
}
