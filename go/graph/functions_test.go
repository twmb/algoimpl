package graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	graph, err := New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	nodes := graph.Nodes()
	// create graph on page 613 of CLRS ed. 3
	nodes = append(nodes, graph.MakeNode("shirt"))
	nodes = append(nodes, graph.MakeNode("tie"))
	nodes = append(nodes, graph.MakeNode("jacket"))
	nodes = append(nodes, graph.MakeNode("belt"))
	nodes = append(nodes, graph.MakeNode("watch"))
	nodes = append(nodes, graph.MakeNode("undershorts"))
	nodes = append(nodes, graph.MakeNode("pants"))
	nodes = append(nodes, graph.MakeNode("shoes"))
	nodes = append(nodes, graph.MakeNode("socks"))
	graph.Connect(nodes[0], nodes[1])
	graph.Connect(nodes[1], nodes[2])
	graph.Connect(nodes[0], nodes[3])
	graph.Connect(nodes[3], nodes[2])
	graph.Connect(nodes[5], nodes[6])
	graph.Connect(nodes[5], nodes[7])
	graph.Connect(nodes[6], nodes[3])
	graph.Connect(nodes[6], nodes[7])
	graph.Connect(nodes[8], nodes[7])
	graph.verify(t)
	wantOrder := make([]*Node, len(graph.Nodes()))
	wantOrder[0] = nodes[8] // socks
	wantOrder[1] = nodes[5] // undershorts
	wantOrder[2] = nodes[6] // pants
	wantOrder[3] = nodes[7] // shoes
	wantOrder[4] = nodes[4] // watch
	wantOrder[5] = nodes[0] // shirt
	wantOrder[6] = nodes[3] // belt
	wantOrder[7] = nodes[1] // tie
	wantOrder[8] = nodes[2] // jacket
	result := TopologicalSort(graph)
	for i := range result {
		if result[i] != wantOrder[i] {
			t.Errorf("index %v in result != wanted, value: %v, want value: %v", i, result[i].Value, wantOrder[i].Value)
		}
	}
}
