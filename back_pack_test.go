package mygomod

// nums = [-1,0,3,5,9,12], target = 9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type backPackParam struct {
	Cap, Num, Ret int
	Wt, Val       []int
}

var backPackParams = []backPackParam{
	{
		Cap: 4,
		Num: 3,
		Wt:  []int{2, 1, 3},
		Val: []int{4, 2, 3},
		Ret: 6,
	},
}

func Test_backPack(t *testing.T) {
	for _, param := range backPackParams {
		t.Run("", func(t *testing.T) {
			ret := backPack(param.Cap, param.Num, param.Wt, param.Val)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
