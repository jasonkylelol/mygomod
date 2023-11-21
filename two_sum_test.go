package mygomod

// input：nums = [2,7,11,15], target = 9
// output：[0,1]
//  nums[0] + nums[1] == 9 ，return [0, 1] 。

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type twoSumParam struct {
	Nums   []int
	Target int
	Ret    []int
}

var twoSumParams = []twoSumParam{
	{
		Nums:   []int{11, 7, 15, 2},
		Target: 9,
		Ret:    []int{1, 3},
	},
}

func Test_twoSum(t *testing.T) {
	for _, param := range twoSumParams {
		t.Run("", func(t *testing.T) {
			ret := twoSum(param.Nums, param.Target)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
