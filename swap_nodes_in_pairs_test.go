package mygomod

import "testing"

func Test_swapPairs(t *testing.T) {
	t.Run("", func(t *testing.T) {
		list := convLinkedList([]int{1, 2, 3, 4})
		node := swapPairs(list)
		printLinkedList(node)
	})
}
