package mygomod

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	t.Run("", func(t *testing.T) {
		list1 := convLinkedList([]int{1, 2, 4, 5, 9})
		list2 := convLinkedList([]int{1, 3, 4, 7, 13})
		// list1 := convLinkedList([]int{})
		// list2 := convLinkedList([]int{0})
		res := mergeTwoLists(list1, list2)
		printLinkedList(res)
	})
}
