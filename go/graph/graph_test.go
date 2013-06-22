package graph

import (
	"testing"
)

func (g *Graph) edges(from, to *Node) bool {
	for _, v := range g.adjacents[from] {
		if v == to {
			return true
		}
	}
	return false
}

func (g *Graph) verify(t *testing.T) {
	// over all the nodes
	for i, node := range g.nodes {
		if node.Index != i {
			t.Errorf("node's graph index %v != actual graph index %v", node.Index, i)
		}
		// over each adjacent node
		for _, adjacentNode := range g.adjacents[node] {
			// check that the graph contains it in the correct position
			if adjacentNode.Index >= len(g.nodes) ||
				g.nodes[adjacentNode.Index] != adjacentNode {
				t.Errorf("adjacent node %v does not belong to the graph", adjacentNode)
			}
			// if the graph is undirected, check that the adjacent node contains the original node back
			if g.kind == 0 {
				if !g.edges(adjacentNode, node) {
					t.Errorf("undirected graph: node %v has adjacent node %v, adjacent node doesn't contain back", node, adjacentNode)
				}
			}
		}
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		InKind    string
		WantKind  int
		WantError string
	}{
		{"directed", 1, ""},
		{"undirected", 0, ""},
		{"", 0, "Unrecognized graph kind"},
	}

	for _, test := range tests {
		got, err := New(test.InKind)
		if err != nil {
			if err.Error() != test.WantError {
				t.Errorf("Received an error %v != wanted error %v", err, test.WantError)
			}
			continue
		}
		if got.kind != test.WantKind {
			t.Errorf("Received wrong kind of graph")
		}
		if len(got.nodes) > 0 {
			t.Errorf("Received new graph has nodes %v, shouldn't", got.nodes)
		}
	}
}

func TestMakeNode(t *testing.T) {
	graph, err := New("undirected")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create undirected graph")
	}
	nodes := make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode().Index] = i
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if nodes[node.Index] != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	nodes = make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode().Index] = i
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if nodes[node.Index] != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
}

func TestConnect(t *testing.T) {
	graph, err := New("undirected")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create undirected graph")
	}
	mapped := make(map[*Node]int, 0)
	for i := 0; i < 10; i++ {
		mapped[graph.MakeNode()] = i
	}
	nodes := graph.nodes
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(nodes[i], nodes[(i+1+j)%len(nodes)])
		}
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if mapped[node] != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	mapped = make(map[*Node]int, 0)
	for i := 0; i < 10; i++ {
		mapped[graph.MakeNode()] = i
	}
	nodes = graph.nodes
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(nodes[i], nodes[(i+1+j)%len(nodes)])
		}
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if mapped[node] != i {
			t.Errorf("Node at index %v = %v, != %v, wrong!", i, mapped[node], i)
		}
	}
}
