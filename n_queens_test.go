package mygomod

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nQueensParam struct {
	N int

	Ret [][]string
}

var nQueensParams = []nQueensParam{
	{
		N: 4,
		Ret: [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		},
	},
}

func Test_nQueens(t *testing.T) {
	for _, param := range nQueensParams {
		t.Run("", func(t *testing.T) {
			ret := solveNQueens(param.N)
			assert.Equal(t, param.Ret, ret)
			for _, solution := range ret {
				for _, row := range solution {
					fmt.Println(row)
				}
				fmt.Println("")
			}
		})
	}
}
