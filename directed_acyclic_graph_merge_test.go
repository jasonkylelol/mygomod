package mygomod

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_parseDAGNodes(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nodes := map[string]*Node{ // this nodes has two tag: one and two
			"A1": {
				Id:     "A1",
				Tag:    "One",
				Childs: []string{"B1", "C1"},
			},
			"B1": {
				Id:     "B1",
				Childs: []string{"D1"},
			},
			"C1": {
				Id:     "C1",
				Childs: []string{"D1"},
			},
			"D1": {
				Id:     "D1",
				Childs: []string{"A2"},
			},
			"A2": {
				Id:     "A2",
				Tag:    "Two",
				Childs: []string{"B2"},
			},
			"B2": {
				Id:     "B2",
				Childs: []string{"C2"},
			},
			"C2": {
				Id:     "C2",
				Childs: []string{"D2"},
			},
			"D2": {
				Id:     "D2",
				Childs: []string{},
			},
		}
		nodeMap := parseDAGNodes(nodes)
		nodeJson, _ := json.Marshal(nodeMap)
		fmt.Println(string(nodeJson))
	})
}
