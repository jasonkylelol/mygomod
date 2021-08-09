package mygomod

// nums = [-1,0,3,5,9,12], target = 9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type binarySearchParam struct {
	Nums   []int
	Target int
	Ret    int
}

var binarySearchParams = []binarySearchParam{
	{
		Nums:   []int{-1, 0, 3, 5, 9, 12},
		Target: 9,
		Ret:    4,
	},
}

func Test_binarySearch(t *testing.T) {
	for _, param := range binarySearchParams {
		t.Run("", func(t *testing.T) {
			ret := binarySearch(param.Nums, param.Target)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
