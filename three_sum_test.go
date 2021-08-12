package mygomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type threeSumParam struct {
	Nums []int
	Ret  [][]int
}

var threeSumParams = []threeSumParam{
	{
		Nums: []int{-1, 0, 1, 2, -1, -4},
		Ret: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
}

func Test_threeSum(t *testing.T) {
	for _, param := range threeSumParams {
		t.Run("", func(t *testing.T) {
			ret := threeSum(param.Nums)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
