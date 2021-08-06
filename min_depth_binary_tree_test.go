package mygomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type minDepthParam struct {
	root *TreeNode
	Ret  int
}

var minDepthParams = []minDepthParam{
	{
		root: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 9,
			},
			Left: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Ret: 2,
	},
}

func Test_minDepth(t *testing.T) {
	for _, param := range minDepthParams {
		t.Run("", func(t *testing.T) {
			ret := minDepth(param.root)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
