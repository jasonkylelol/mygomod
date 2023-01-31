package mygomod

import (
	"fmt"
	"testing"
)

func Test_sliceReflect(t *testing.T) {
	anyList := []string{
		"[123.675,116.28,103.53]",
		"[11,22,3]",
		"[one,two_dup-v2.1,app-det_v1.0.yaml]",
		"33",
		"400",
		"400",
		"33.25",
		"true",
		"False",
		"rgb",
		"../class_names.txt",
		"111,222,333",
		"[aaaa true",
		"[]",
	}
	for _, nany := range anyList {
		t.Run(nany, func(t *testing.T) {
			fmt.Println("output:", convRunParamStrValue(nany))
			fmt.Println()
		})
	}
}
