// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
	"fmt"
)

type GraphType int

const (
	Undirected = iota
	Directed
)

// An adjacency slice representation of a graph. Can be directed or undirected.
type Graph struct {
	nodes []*node
	edges map[*node][]Edge
	kind  int // 1 for directed, 0 otherwise
}

func (g *Graph) String() string {
	rVal := "g->{\n"
	for _, node := range g.nodes {
		rVal += "\t" + node.String() + "->{"
		for _, adj := range g.edges[node] {
			rVal += adj.Start.node.String() + "->" + adj.End.node.String() + ","
		}
		rVal += "}\n"
	}
	rVal += "}"
	return rVal
}

type node struct {
	graphIndex int
	state      int   // used for metadata
	data       int   // also used for metadata
	parent     *node // also used for metadata
	container  Node  // who holds me
}

// Node connects to a node on the graph. It can safely be used in maps.
type Node struct {
	// In an effort to prevent access to the actual graph
	// and so that the Node type can be used in a map while
	// the graph changes metadata, the Node type encapsulates
	// a pointer to the actual node data.
	node *node
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.graphIndex)
}

type Edge struct {
	Weight *int
	Start  Node
	End    Node
	Kind   GraphType
}

func (e Edge) String() string {
	rVal := fmt.Sprintf("{%v", *e.Weight)
	rVal += " " + e.Start.node.String() + "->"
	rVal += e.End.node.String() + "}"
	return rVal
}

// In case the user abused newG := &Graph{}...
func (g *Graph) lazyInit() {
	if g.nodes == nil {
		g.nodes = []*node{}
		g.edges = make(map[*node][]Edge, 0)
	}
}

// Creates and returns an empty graph.
// If kind is "directed", returns a directed graph.
// If kind is "undirected", this function will return an undirected graph.
// Otherwise, this will return nil and an error.
func New(kind string) (*Graph, error) {
	switch kind {
	case "directed":
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge), kind: Directed}, nil
	case "undirected":
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge)}, nil
	default:
		return nil, errors.New("Unrecognized graph kind")
	}
}

// Creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	g.lazyInit()
	newNode := &node{graphIndex: len(g.nodes)}
	newNode.container = Node{node: newNode}
	g.nodes = append(g.nodes, newNode)
	return newNode.container
}

// Creates and an edge and returns a pointer to a copy of the edge.
// The return value will be nil if either of the nodes do not belong
// to the graph.
//
// Calling connect multiple times on the same nodes will not
// make multiple edges; the same edge will be returned on each call.
func (g *Graph) Connect(from, to Node) *Edge {
	if from.node.graphIndex >= len(g.nodes) || g.nodes[from.node.graphIndex] != from.node {
		return (*Edge)(nil)
	}
	if to.node.graphIndex >= len(g.nodes) || g.nodes[to.node.graphIndex] != to.node {
		return (*Edge)(nil)
	}
	for _, edge := range g.edges[from.node] { // check if edge already exists
		if edge.End == to || edge.Start == to && edge.End == from {
			copyEdge := edge
			return &copyEdge
		}
	}
	newEdge := Edge{Weight: new(int), Start: from, End: to}
	if g.kind == Directed {
		newEdge.Kind = Directed
	}
	g.edges[from.node] = append(g.edges[from.node], newEdge)
	if g.kind == Undirected { // for book keeping
		g.edges[to.node] = append(g.edges[to.node], newEdge)
	}
	copyEdge := newEdge
	return &copyEdge
}
