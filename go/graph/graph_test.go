package graph

import (
	"testing"
)

func (g *Graph) edgeBack(from, to *node) bool {
	for _, v := range g.edges[from] {
		if v.End.node == to || v.Start.node == to && v.End.node == from {
			return true
		}
	}
	return false
}

func (g *Graph) verify(t *testing.T) {
	// over all the nodes
	for i, node := range g.nodes {
		if node.graphIndex != i {
			t.Errorf("node's graph index %v != actual graph index %v", node.graphIndex, i)
		}
		// over each adjacent node
		for _, edge := range g.edges[node] {
			// check that the graph contains it in the correct position
			if edge.End.node.graphIndex >= len(g.nodes) ||
				g.nodes[edge.End.node.graphIndex] != edge.End.node {
				t.Errorf("adjacent node %v does not belong to the graph", edge.End)
			}
			// if the graph is undirected, check that the adjacent node contains the original node back
			if g.kind == Undirected {
				if !g.edgeBack(edge.End.node, node) {
					t.Errorf("undirected graph: node %v has adjacent node %v, adjacent node doesn't contain back", node, edge.End)
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
	nodes := make(map[Node]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode()] = i
	}
	graph.verify(t)
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	nodes = make(map[Node]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode()] = i
	}
	graph.verify(t)
}

func TestConnect(t *testing.T) {
	graph, err := New("undirected")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create undirected graph")
	}
	mapped := make(map[int]Node, 0)
	for i := 0; i < 10; i++ {
		mapped[i] = graph.MakeNode()
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(mapped[i], mapped[(i+1+j)%10])
		}
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if mapped[i].node != node {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	mapped = make(map[int]Node, 0)
	for i := 0; i < 10; i++ {
		mapped[i] = graph.MakeNode()
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(mapped[i], mapped[(i+1+j)%10])
		}
	}
	graph.verify(t)
	for i, node := range graph.nodes {
		if mapped[i].node != node {
			t.Errorf("Node at index %v = %v, != %v, wrong!", i, mapped[i], node)
		}
	}
}
