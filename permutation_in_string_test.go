package mygomod

// s1 = "ab" s2 = "eidbaooo"
// true

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type permutationInStringParam struct {
	S1, S2 string

	Ret bool
}

var permutationInStringParams = []permutationInStringParam{
	{
		S1:  "ab",
		S2:  "eidbaooo",
		Ret: true,
	},
}

func Test_permutationInString(t *testing.T) {
	for _, param := range permutationInStringParams {
		t.Run("", func(t *testing.T) {
			ret := checkInclusion(param.S1, param.S2)
			assert.Equal(t, param.Ret, ret)
		})
	}
}
