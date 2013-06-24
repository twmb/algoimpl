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
		// over each edge
		for _, edge := range g.edges[node] {

			// check that the graph contains it in the correct position
			if edge.End.node.graphIndex >= len(g.nodes) {
				t.Errorf("adjacent node end graph index %v >= len(g.nodes)%v", edge.End.node.graphIndex, len(g.nodes))
			}
			if g.nodes[edge.End.node.graphIndex] != edge.End.node {
				t.Errorf("adjacent node %p does not belong to the graph on edge %v: should be %p", edge.End.node, edge, g.nodes[edge.End.node.graphIndex])
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
		InKind   GraphType
		WantKind GraphType
	}{
		{Directed, 1},
		{Undirected, 0},
		{3, 0},
	}

	for _, test := range tests {
		got := New(test.InKind)
		if got.kind != test.WantKind {
			t.Errorf("Received wrong kind of graph")
		}
		if len(got.nodes) > 0 {
			t.Errorf("Received new graph has nodes %v, shouldn't", got.nodes)
		}
	}
}

func TestMakeNode(t *testing.T) {
	graph := New(Undirected)
	nodes := make(map[Node]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode()] = i
	}
	graph.verify(t)
	graph = New(Directed)
	nodes = make(map[Node]int, 0)
	for i := 0; i < 10; i++ {
		nodes[graph.MakeNode()] = i
	}
	graph.verify(t)
}

func TestRemoveNode(t *testing.T) {
	g := New(Undirected)
	nodes := make([]Node, 2)
	nodes[0] = g.MakeNode()
	nodes[1] = g.MakeNode()
	g.Connect(nodes[0], nodes[0])
	g.Connect(nodes[1], nodes[0])
	g.Connect(nodes[0], nodes[1])
	g.Connect(nodes[1], nodes[1])
	g.verify(t)
	g.RemoveNode(&nodes[1])
	g.verify(t)
	g.RemoveNode(&nodes[1])
	g.verify(t)
	g.RemoveNode(&nodes[0])
	g.verify(t)
	nodes = make([]Node, 10)
	g = New(Directed)
	for i := 0; i < 10; i++ {
		nodes[i] = g.MakeNode()
	}
	// connect every node to every node
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if g.Connect(nodes[i], nodes[j]) != nil {
				t.Errorf("could not connect %v, %v", i, j)
			}
		}
	}
	g.verify(t)
	g.RemoveNode(&nodes[0])
	g.verify(t)
	if nodes[0].node != nil {
		t.Errorf("Node still has reference to node in graph")
	}
	g.RemoveNode(&nodes[9])
	g.verify(t)
	g.RemoveNode(&nodes[9])
	g.verify(t)
	g.RemoveNode(&nodes[0])
	g.verify(t)
	g.RemoveNode(&nodes[1])
	g.verify(t)
	g.RemoveNode(&nodes[2])
	g.verify(t)
	g.RemoveNode(&nodes[3])
	g.verify(t)
	g.RemoveNode(&nodes[4])
	g.verify(t)
	g.RemoveNode(&nodes[5])
	g.verify(t)
	g.RemoveNode(&nodes[6])
	g.verify(t)
	g.RemoveNode(&nodes[7])
	g.verify(t)
	g.RemoveNode(&nodes[8])
	g.verify(t)
	g.RemoveNode(&nodes[9])
	g.verify(t)
}

func TestConnect(t *testing.T) {
	graph := New(Undirected)
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
	graph = New(Directed)
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

func TestUnconnect(t *testing.T) {
	g := New(Undirected)
	nodes := make([]Node, 2)
	nodes[0] = g.MakeNode()
	nodes[1] = g.MakeNode()
	g.Connect(nodes[0], nodes[0])
	g.Connect(nodes[1], nodes[0])
	g.Connect(nodes[0], nodes[1])
	g.Connect(nodes[1], nodes[1])
	g.verify(t)
	g.Unconnect(nodes[0], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[0], nodes[1])
	g.verify(t)
	g.Unconnect(nodes[1], nodes[1])
	g.verify(t)
	nodes = make([]Node, 10)
	g = New(Directed)
	for i := 0; i < 10; i++ {
		nodes[i] = g.MakeNode()
	}
	// connect every node to every node
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if g.Connect(nodes[i], nodes[j]) != nil {
				t.Errorf("could not connect %v, %v", i, j)
			}
		}
	}
	g.verify(t)
	g.Unconnect(nodes[5], nodes[4])
	g.verify(t)
	g.Unconnect(nodes[9], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[9], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[0], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[1], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[2], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[3], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[4], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[5], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[6], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[7], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[8], nodes[0])
	g.verify(t)
	g.Unconnect(nodes[9], nodes[0])
	g.verify(t)
}
