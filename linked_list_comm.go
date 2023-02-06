package mygomod

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func convLinkedList(inputs []int) *ListNode {
	if len(inputs) == 0 {
		return nil
	}
	node := &ListNode{Val: inputs[0]}
	head := node
	for idx := 1; idx < len(inputs); idx++ {
		node.Next = &ListNode{Val: inputs[idx]}
		node = node.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	var ss strings.Builder
	ss.WriteString("[")
	idx := 0
	for head != nil {
		if idx != 0 {
			ss.WriteString(",")
		}
		idx++
		ss.WriteString(fmt.Sprintf("%v", head.Val))
		head = head.Next
	}
	ss.WriteString("]")
	fmt.Println(ss.String())
}
