// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
	"fmt"
)

// An adjacency slice representation of a graph. Can be directed or undirected.
type Graph struct {
	nodes []*Node
	kind  int // 1 for directed, 0 otherwise
}

// Prints in the following format:
//    g->{
//      node->{node, node, node,}
//      node->{node, node, node,}
//    }
func (g *Graph) String() string {
	rVal := "g->{\n"
	for i := range g.nodes {
		rVal += "\t" + g.nodes[i].String() + "\n"
	}
	rVal += "}\n"
	return rVal
}

// Contains a value on the graph and
// other metadata needed for functions.
type Node struct {
	Value      interface{}
	adjacent   []*Node
	graphIndex int
	state      int   // used for sort / search / other functions as metadata
	parent     *Node // also used as metadata
}

// Prints in the following format:
//    value->{value, value, value,}
// Where -> represents an edges from a node to other nodes.
func (n *Node) String() string {
	rVal := ""
	rVal += fmt.Sprint(n.Value) + "->{"
	for adj := range n.adjacent {
		rVal += fmt.Sprint(n.adjacent[adj].Value) + ", "
	}
	rVal += "}"
	return rVal
}

// Incase the user abused newG := &Graph{}...
func (g *Graph) lazyInit() {
	if g.nodes == nil {
		g.nodes = []*Node{}
	}
}

// Creates and returns an empty graph.
// If kind is "directed", returns a directed graph.
// If kind is "undirected", this function will return an undirected graph.
// Otherwise, this will return nil and an error.
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

// Creates an edge between two nodes in a graph.
// If the graph is undirected, this function also connects the to node to the from node.
// Returns an error if either of the nodes do not belong to the graph.
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
