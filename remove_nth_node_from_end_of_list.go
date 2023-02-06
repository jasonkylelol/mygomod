package mygomod

// Given the head of a linked list, remove the n-th node from the end of the list and return its head.
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	steps := 0
	for fast != nil {
		fast = fast.Next
		if steps > n {
			slow = slow.Next
		}
		steps++
	}
	if steps == n {
		head = head.Next
	} else if steps < n && slow == head {
		return head
	} else {
		slow.Next = slow.Next.Next
	}
	return head
}
