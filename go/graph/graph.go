// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
	"fmt"
)

// Directed or undirected.
type GraphType int

const (
	Undirected GraphType = iota
	Directed
)

// An adjacency slice representation of a graph. Can be directed or undirected.
type Graph struct {
	nodes []*node
	edges map[*node][]Edge
	kind  GraphType // 1 for directed, 0 otherwise
}

// Prints in the format
//     g->{
//         0->{0->0, 1->0,}
//         1->{2->1, 1->0,}
//         2->{2->1,}
//         3->{}
//     }
// Where numbers represent graph indices and -> points to nodes indices they connect to.
// If the graph is undirected, one edge (i.e., 2->1) is represented on both nodes.
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

func (n *node) String() string {
	return fmt.Sprintf("%v", n.graphIndex)
}

// Node connects to a backing node on the graph. It can safely be used in maps.
type Node struct {
	// In an effort to prevent access to the actual graph
	// and so that the Node type can be used in a map while
	// the graph changes metadata, the Node type encapsulates
	// a pointer to the actual node data.
	node *node
}

// An edge connects two Nodes in a graph. The weight can be modified and
// used for functions that rely on weights.
//
// In an undirected graph, the start of an edge and end of an edge
// is represented once in the graph: if you connect A to B
// and use the Remove function to remove B, the returned edge will
// have a Start of A and an End of B.
type Edge struct {
	Weight *int
	Start  Node
	End    Node
	Kind   GraphType
}

// Prints in the following format:
//     {5 2->3}
// Where 5 is the weight and 2->3 is an edge.
func (e Edge) String() string {
	rVal := fmt.Sprintf("{%v", *e.Weight)
	rVal += " " + e.Start.node.String() + "->"
	rVal += e.End.node.String() + "}"
	return rVal
}

// Creates and returns an empty graph. This function must be called before nodes can be connected.
// If kind is Directed, returns a directed graph.
// If kind is Undirected, this function will return an undirected graph.
// Otherwise, this will return nil and an error.
func New(kind GraphType) (*Graph, error) {
	switch kind {
	case Directed:
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge), kind: Directed}, nil
	case Undirected:
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge)}, nil
	default:
		return nil, errors.New("Unrecognized graph kind")
	}
}

// Creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	newNode := &node{graphIndex: len(g.nodes)}
	newNode.container = Node{node: newNode}
	g.nodes = append(g.nodes, newNode)
	return newNode.container
}

// Removes a node from the graph and all edges connected to it
// and nil's all connections on the node for garbage collection.
// Because the node that `remove` points to will be nilled, if
// the node is used in a map, you can no longer access that element
// in the map. Delete the map index first.
// Has O(V+E) time complexity.
func (g *Graph) RemoveNode(remove *Node) {
	// O(V)
	if remove.node == nil {
		return
	}
	nodeExists := false
	for _, node := range g.nodes {
		edges := g.edges[node]
		// O(E)
		swapIndex := 0 // index edge to remove is at: swap this with end
		needSwap := false
		for edgeI, edge := range edges {
			if edge.Start == *remove || edge.End == *remove {
				nodeExists = true
				edge.Start.node = nil
				edge.End.node = nil
				swapIndex = edgeI
				needSwap = true
			}
		}
		if needSwap {
			edges[swapIndex], edges[len(edges)-1] = edges[len(edges)-1], edges[swapIndex]
			edges = edges[:len(edges)-1]
		}
		g.edges[node] = edges
		if node.graphIndex > remove.node.graphIndex {
			node.graphIndex--
		}
	}
	if remove.node.graphIndex < len(g.nodes)-1 {
		copy(g.nodes[remove.node.graphIndex:], g.nodes[remove.node.graphIndex+1:])
	}
	if nodeExists {
		g.nodes = g.nodes[:len(g.nodes)-1]
	}
	remove.node.parent = nil
	remove.node = nil
}

// Creates an edge and returns a pointer to a copy of the edge.
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
	if g.kind == Undirected && to != from { // for book keeping
		g.edges[to.node] = append(g.edges[to.node], newEdge)
	}
	copyEdge := newEdge
	return &copyEdge
}
