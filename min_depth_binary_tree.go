package mygomod

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%v", t.Val)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	step := 1
	que := []*TreeNode{root}
	for len(que) > 0 {
		size := len(que)
		for idx := 0; idx < size; idx++ {
			cur := que[0]
			que = que[1:]
			fmt.Println(idx, ":", size, ":", cur.Val)
			if cur.Left == nil && cur.Right == nil {
				return step
			}
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
			fmt.Println(que)
		}
		step++
	}
	return step
}
