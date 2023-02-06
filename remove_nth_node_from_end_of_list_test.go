package mygomod

import "testing"

func Test_removeNthFromEnd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		node := convLinkedList([]int{1, 2, 3, 4, 5})
		res := removeNthFromEnd(node, 2)
		printLinkedList(res)
	})
}
