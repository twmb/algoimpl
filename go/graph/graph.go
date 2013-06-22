// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
	"fmt"
)

// An adjacency slice representation of a graph. Can be directed or undirected.
type Graph struct {
	nodes     []*Node
	adjacents map[*Node][]*Node
	kind      int // 1 for directed, 0 otherwise
}

func (g *Graph) String() string {
	rVal := "g->{\n"
	for _, node := range g.nodes {
		rVal += "\t" + node.String() + "->{"
		for _, adj := range g.adjacents[node] {
			rVal += adj.String() + ","
		}
		rVal += "}\n"
	}
	rVal += "}"
	return rVal
}

type Node struct {
	Index  int
	state  int   // used for metadata
	parent *Node // also used for metadata
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Index)
}

// Incase the user abused newG := &Graph{}...
func (g *Graph) lazyInit() {
	if g.nodes == nil {
		g.nodes = []*Node{}
		g.adjacents = make(map[*Node][]*Node, 0)
	}
}

// Creates and returns an empty graph.
// If kind is "directed", returns a directed graph.
// If kind is "undirected", this function will return an undirected graph.
// Otherwise, this will return nil and an error.
func New(kind string) (*Graph, error) {
	switch kind {
	case "directed":
		return &Graph{nodes: []*Node{}, adjacents: make(map[*Node][]*Node), kind: 1}, nil
	case "undirected":
		return &Graph{nodes: []*Node{}, adjacents: make(map[*Node][]*Node)}, nil
	default:
		return nil, errors.New("Unrecognized graph kind")
	}
}

// Creates a node, adds it to the graph and returns the new node.
// A possibly way to manage what nodes are mapped to what values is to maintain a
//     map[int]Value
// on the graph caller side, with the int being the node.Index
func (g *Graph) MakeNode() *Node {
	g.lazyInit()
	newNode := &Node{Index: len(g.nodes)}
	g.nodes = append(g.nodes, newNode)
	return newNode
}

// Creates an edge between two nodes in a graph.
// If the graph is undirected, this function also connects the to node to the from node.
// Example usage:
//     graph.Connect(node1, node2)
// Returns an error if either of the nodes do not belong to the graph.
func (g *Graph) Connect(from, to *Node) error {
	if from.Index >= len(g.nodes) || g.nodes[from.Index] != from {
		return errors.New("from node in connect call does not belong to the graph")
	}
	if to.Index >= len(g.nodes) || g.nodes[to.Index] != to {
		return errors.New("to node in connect call does not belong to the graph")
	}
	g.adjacents[from] = append(g.adjacents[from], to)
	if g.kind == 0 { // undirected graph
		g.adjacents[to] = append(g.adjacents[to], from)
	}
	return nil
}
