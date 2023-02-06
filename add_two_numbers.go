package mygomod

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	var head *ListNode
	upgrade := 0
	for l1 != nil || l2 != nil {
		sum := upgrade
		if l1 != nil && l2 != nil {
			sum += l1.Val + l2.Val
		} else if l2 != nil {
			sum += l2.Val
		} else if l1 != nil {
			sum += l1.Val
		} else {
			break
		}
		node.Next = &ListNode{Val: sum % 10}
		if head == nil {
			head = node.Next
		}
		upgrade = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		node = node.Next
	}
	if upgrade > 0 {
		node.Next = &ListNode{Val: upgrade}
	}
	return head
}
