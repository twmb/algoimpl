package graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	graph := New(Directed)
	nodes := make([]Node, 0)
	// create graph on page 613 of CLRS ed. 3
	nodes = append(nodes, graph.MakeNode()) // shirt
	nodes = append(nodes, graph.MakeNode()) // tie
	nodes = append(nodes, graph.MakeNode()) // jacket
	nodes = append(nodes, graph.MakeNode()) // belt
	nodes = append(nodes, graph.MakeNode()) // watch
	nodes = append(nodes, graph.MakeNode()) // undershorts
	nodes = append(nodes, graph.MakeNode()) // pants
	nodes = append(nodes, graph.MakeNode()) // shoes
	nodes = append(nodes, graph.MakeNode()) // socks
	graph.MakeEdge(nodes[0], nodes[1])
	graph.MakeEdge(nodes[1], nodes[2])
	graph.MakeEdge(nodes[0], nodes[3])
	graph.MakeEdge(nodes[3], nodes[2])
	graph.MakeEdge(nodes[5], nodes[6])
	graph.MakeEdge(nodes[5], nodes[7])
	graph.MakeEdge(nodes[6], nodes[3])
	graph.MakeEdge(nodes[6], nodes[7])
	graph.MakeEdge(nodes[8], nodes[7])
	graph.verify(t)
	wantOrder := make([]Node, len(graph.nodes))
	wantOrder[0] = nodes[8] // socks
	wantOrder[1] = nodes[5] // undershorts
	wantOrder[2] = nodes[6] // pants
	wantOrder[3] = nodes[7] // shoes
	wantOrder[4] = nodes[4] // watch
	wantOrder[5] = nodes[0] // shirt
	wantOrder[6] = nodes[3] // belt
	wantOrder[7] = nodes[1] // tie
	wantOrder[8] = nodes[2] // jacket
	result := graph.TopologicalSort()
	for i := range result {
		if result[i] != wantOrder[i] {
			t.Errorf("index %v in result != wanted, value: %v, want value: %v", i, result[i], wantOrder[i])
		}
	}
}

func TestStronglyConnectedComponents(t *testing.T) {
	graph := New(Directed)
	nodes := make([]Node, 0)
	// create SCC graph on page 616 of CLRS ed 3
	nodes = append(nodes, graph.MakeNode()) //0, c
	nodes = append(nodes, graph.MakeNode()) //1, g
	nodes = append(nodes, graph.MakeNode()) //2, f
	nodes = append(nodes, graph.MakeNode()) //3, h
	nodes = append(nodes, graph.MakeNode()) //4, d
	nodes = append(nodes, graph.MakeNode()) //5, b
	nodes = append(nodes, graph.MakeNode()) //6, e
	nodes = append(nodes, graph.MakeNode()) //7, a
	graph.MakeEdge(nodes[0], nodes[1])
	graph.MakeEdge(nodes[0], nodes[4])
	graph.MakeEdge(nodes[1], nodes[2])
	graph.MakeEdge(nodes[1], nodes[3])
	graph.MakeEdge(nodes[2], nodes[1])
	graph.MakeEdge(nodes[3], nodes[3])
	graph.MakeEdge(nodes[4], nodes[3])
	graph.MakeEdge(nodes[4], nodes[0])
	graph.MakeEdge(nodes[5], nodes[6])
	graph.MakeEdge(nodes[5], nodes[0])
	graph.MakeEdge(nodes[5], nodes[2])
	graph.MakeEdge(nodes[6], nodes[2])
	graph.MakeEdge(nodes[6], nodes[7])
	graph.MakeEdge(nodes[7], nodes[5])
	graph.verify(t)
	want := make([][]Node, 4)
	want[0] = make([]Node, 3)
	want[1] = make([]Node, 2)
	want[2] = make([]Node, 2)
	want[3] = make([]Node, 1)
	want[0][0] = nodes[6]
	want[0][1] = nodes[7]
	want[0][2] = nodes[5]
	want[1][0] = nodes[0]
	want[1][1] = nodes[4]
	want[2][0] = nodes[2]
	want[2][1] = nodes[1]
	want[3][0] = nodes[3]
	components := graph.StronglyConnectedComponents()
	for j := range components {
		for i := range want[j] {
			if !componentContains(components[j], want[j][i]) {
				t.Errorf("component slice %v does not contain want node %v", components[j], want[j][i])
			}
		}
	}
}

func componentContains(component []Node, node Node) bool {
	for i := range component {
		if component[i].node.index == node.node.index { // for SCC, the nodes will be reversed but the indices will be the same
			return true
		}
	}
	return false
}
