package mygomod

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	t.Run("", func(t *testing.T) {
		node := addTwoNumbers(convLinkedList([]int{2, 4, 3}), convLinkedList([]int{5, 6, 7}))
		printLinkedList(node)
	})
}
