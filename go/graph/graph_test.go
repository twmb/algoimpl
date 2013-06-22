package graph

import (
	"testing"
)

func (n *Node) contains(other *Node) bool {
	for _, v := range n.adjacent {
		if v == other {
			return true
		}
	}
	return false
}

func valueSliceContains(slice []*Node, key interface{}) bool {
	for i := range slice {
		if slice[i].Value == key {
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
		for _, adjacentNode := range node.adjacent {
			// check that the graph contains it in the correct position
			if adjacentNode.graphIndex >= len(g.nodes) ||
				g.nodes[adjacentNode.graphIndex] != adjacentNode {
				t.Errorf("adjacent node %v does not belong to the graph", adjacentNode)
			}
			// if the graph is undirected, check that the adjacent node contains the original node back
			if g.kind == 0 {
				if !adjacentNode.contains(node) {
					t.Errorf("undirected graph: node %v has adjacent node %v, adjacent node doesn't contain back", node.Value, adjacentNode.Value)
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
	for i := 0; i < 10; i++ {
		graph.MakeNode(i)
	}
	graph.verify(t)
	for i, node := range graph.Nodes() {
		if node.Value != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	for i := 0; i < 10; i++ {
		graph.MakeNode(i)
	}
	graph.verify(t)
	for i, node := range graph.Nodes() {
		if node.Value != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
}

func TestConnect(t *testing.T) {
	graph, err := New("undirected")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create undirected graph")
	}
	for i := 0; i < 10; i++ {
		graph.MakeNode(i)
	}
	nodes := graph.Nodes()
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(nodes[i], nodes[(i+1+j)%len(nodes)])
		}
	}
	graph.verify(t)
	for i, node := range graph.Nodes() {
		if node.Value != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
	graph, err = New("directed")
	if err != nil {
		t.Errorf("TestMakeNode: unable to create directed graph")
	}
	for i := 0; i < 10; i++ {
		graph.MakeNode(i)
	}
	nodes = graph.Nodes()
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			graph.Connect(nodes[i], nodes[(i+1+j)%len(nodes)])
		}
	}
	graph.verify(t)
	for i, node := range graph.Nodes() {
		if node.Value != i {
			t.Errorf("Node at index %v != %v, wrong!", i, i)
		}
	}
}
