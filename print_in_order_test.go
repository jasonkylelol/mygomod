package mygomod

// s1 = "ab" s2 = "eidbaooo"
// true

import (
	"testing"
)

type printInOrderParam struct {
}

var printInOrderParams = []printInOrderParam{
	{},
}

func Test_printInOrder(t *testing.T) {
	for _, param := range printInOrderParams {
		_ = param
		t.Run("", func(t *testing.T) {
			printInOrder()
		})
	}
}
