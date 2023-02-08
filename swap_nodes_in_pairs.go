package mygomod

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	node := head.Next
	head_front := &ListNode{}
	for head != nil {
		front := head
		back := head.Next
		if back == nil {
			break
		}
		// fmt.Println(front.Val, back.Val)
		head_front.Next = back
		dummy := back.Next
		front.Next = dummy
		back.Next = front
		head = dummy
		head_front = front
	}
	return node
}
