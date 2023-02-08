package mygomod

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else if list2 == nil {
			head.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				head.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				head.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
		}
		head = head.Next
	}
	return node.Next
}
