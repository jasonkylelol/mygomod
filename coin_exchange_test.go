package mygomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type coinChangeParam struct {
	Coins  []int
	Amount int
	Ret    int
}

var coinChangeParams = []coinChangeParam{
	{
		Coins:  []int{1, 2, 5},
		Amount: 11,
		Ret:    3,
	},
	{
		Coins:  []int{2},
		Amount: 9,
		Ret:    -1,
	},
}

func Test_coinChange(t *testing.T) {
	for _, param := range coinChangeParams {
		t.Run("", func(t *testing.T) {
			ret := coinChange(param.Coins, param.Amount)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
