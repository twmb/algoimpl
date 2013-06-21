// Implements an adjacency list graph as a slice of generic nodes
package graph

import (
	"errors"
)

type Graph struct {
	nodes []*Node
	kind  int // 1 for directed, 0 otherwise
}

type Node struct {
	Value      interface{}
	adjacent   []*Node
	graphIndex int
}

func (g *Graph) lazyInit() {
	if g.nodes == nil {
		g.nodes = []*Node{}
	}
}

// Creates and returns an empty graph.
// If kind is "directed", returns a directed graph.
// If kind is "undirected", this function will return an undirected graph.
// Otherwise, this will return nil and an error.
// Otherwise, returns an undirected graph.
func New(kind string) (*Graph, error) {
	switch kind {
	case "directed":
		return &Graph{nodes: []*Node{}, kind: 1}, nil
	case "undirected":
		return &Graph{nodes: []*Node{}}, nil
	default:
		return nil, errors.New("Unrecognized graph kind")
	}
}

// Creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode(v interface{}) *Node {
	g.lazyInit()
	newNode := &Node{Value: v, adjacent: []*Node{}, graphIndex: len(g.nodes)}
	g.nodes = append(g.nodes, newNode)
	return newNode
}

// Returns the slice of pointers to the graph for iterating over.
// This package assumes that individual indices will not be modified
// inappropriately. If they are, then the adjacency list structure
// will not hold.
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// Creates an edge between two nodes in a graph.
// If the graph is undirected, this function also connects the to node to the from node.
func (g *Graph) Connect(from, to *Node) error {
	if from.graphIndex >= len(g.nodes) || g.nodes[from.graphIndex] != from {
		return errors.New("from node in connect call does not belong to the graph")
	}
	if to.graphIndex >= len(g.nodes) || g.nodes[to.graphIndex] != to {
		return errors.New("to node in connect call does not belong to the graph")
	}
	from.adjacent = append(from.adjacent, to)
	if g.kind == 0 { // undirected graph
		to.adjacent = append(to.adjacent, from)
	}
	return nil
}

// Returns the slice of pointers to adjacent nodes.
// This package assumes that individual indices will not be modified
// inappropriately. If they are, then the adjacency list structure
// will not hold.
func (n *Node) Adjacent() []*Node {
	return n.adjacent
}
