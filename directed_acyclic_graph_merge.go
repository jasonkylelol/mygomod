package mygomod

import (
	"fmt"
)

type Node struct {
	Id     string   `json:"id"`
	Tag    string   `json:"-"` // only root node cotains tag
	Childs []string `json:"-"`
}

// given nodes map with nodeId as key, return map contains nodes group by tag
// which is directed acyclic graph
func parseDAGNodes(inputs map[string]*Node) map[string][]*Node {
	nodeMap := map[string][]*Node{}
	if len(inputs) > 0 {
		traverDAGNodes("A1", "", inputs, nodeMap)
	}
	return nodeMap
}

func traverDAGNodes(nodeId, tag string, inputs map[string]*Node, nodeMap map[string][]*Node) {
	node := inputs[nodeId]
	if node.Tag != "" {
		tag = node.Tag
	}
	fmt.Println(node.Id, node.Childs)
	_, ok := nodeMap[tag]
	if ok {
		found := false
		for _, n := range nodeMap[tag] {
			if n.Id == node.Id {
				found = true
				break
			}
		}
		if !found {
			nodeMap[tag] = append(nodeMap[tag], node)
		}
	} else {
		nodeMap[tag] = []*Node{node}
	}
	for _, child := range node.Childs {
		traverDAGNodes(child, tag, inputs, nodeMap)
	}
}
