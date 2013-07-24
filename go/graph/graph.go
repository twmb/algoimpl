// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
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
	kind  GraphType // 1 for directed, 0 otherwise
}

type node struct {
	edges         []edge
	reversedEdges []edge
	index         int
	state         int   // used for metadata
	data          int   // also used for metadata
	parent        *node // also used for metadata
	container     Node  // who holds me
}

// Node connects to a backing node on the graph. It can safely be used in maps.
type Node struct {
	// In an effort to prevent access to the actual graph
	// and so that the Node type can be used in a map while
	// the graph changes metadata, the Node type encapsulates
	// a pointer to the actual node data.
	node *node
	// Value can be used to store information on the caller side.
	// Its use is optional. See the Topological Sort example for
	// a reason on why to use this pointer.
	// The reason it is a pointer is so that graph function calls
	// can test for equality on Nodes. The pointer wont change,
	// the value it points to will. If the pointer is explicitly changed,
	// graph functions that use Nodes will cease to work.
	Value *interface{}
}

type edge struct {
	weight int
	end    *node
}

// An edge connects two Nodes in a graph. To modify the weight, use
// the CreateEdgeWeight function. Any local modifications will
// not be seen in the graph.
//
// In an undirected graph, the start of an edge and end of an edge
// is represented once in the graph.
type Edge struct {
	Weight int
	Start  Node
	End    Node
	Kind   GraphType
}

// Creates and returns an empty graph. This function must be called before nodes can be connected.
// If kind is Directed, returns a directed graph.
// If kind is Undirected, this function will return an undirected graph.
// If kind is anything else, this function will return an undirected graph by default.
func New(kind GraphType) *Graph {
	g := &Graph{}
	if kind == Directed {
		g.kind = Directed
	}
	return g
}

// Creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	newNode := &node{index: len(g.nodes)}
	newNode.container = Node{node: newNode, Value: new(interface{})}
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
	// remove all edges that connect from a different node to this one
	for _, node := range g.nodes {
		if node == remove.node {
			// clear memory for everything the removing node contains
			for _, edge := range node.edges {
				edge.end = nil
			}
			for _, edge := range node.reversedEdges {
				edge.end = nil
			}
			continue
		}
		// O(E)
		swapIndex := -1 // index that the edge-to-remove is at: swap this with element at end of slice
		for edgeI, edge := range node.edges {
			if edge.end == remove.node {
				nodeExists = true
				edge.end = nil
				swapIndex = edgeI
			}
		}
		if swapIndex > -1 {
			node.edges[swapIndex], node.edges[len(node.edges)-1] = node.edges[len(node.edges)-1], node.edges[swapIndex]
			node.edges = node.edges[:len(node.edges)-1]
		}
		// now deal with possible reversed edge
		swapIndex = -1
		for edgeI, edge := range node.reversedEdges {
			if edge.end == remove.node {
				nodeExists = true
				edge.end = nil
				swapIndex = edgeI
			}
		}
		if swapIndex > -1 {
			node.reversedEdges[swapIndex], node.reversedEdges[len(node.reversedEdges)-1] = node.reversedEdges[len(node.reversedEdges)-1], node.reversedEdges[swapIndex]
			node.reversedEdges = node.reversedEdges[:len(node.reversedEdges)-1]
		}
		if node.index > remove.node.index {
			node.index--
		}
	}
	if remove.node.index < len(g.nodes)-1 {
		copy(g.nodes[remove.node.index:], g.nodes[remove.node.index+1:])
	}
	if nodeExists {
		g.nodes[len(g.nodes)-1] = nil // garbage collect
		g.nodes = g.nodes[:len(g.nodes)-1]
	}
	remove.node.parent = nil
	remove.node = nil
}

// This function calls CreateEdgeWeight with a weight of 0
// and returns an error if either of the nodes do not
// belong in the graph.
// Calling CreateEdge multiple times on the same nodes will not
// make multiple edges.
//
// Runs in O(E) time, where E is the number of edges coming out
// of the from node (and to node if the graph is undirected).
// If the graph is sparse, which it should be if you are using
// an adjacency list graph structure like this package provides,
// then connecting two nodes will be near O(1) time. For the exact
// specifics on the performance of this function, click on it and
// read the function.
func (g *Graph) CreateEdge(from, to Node) error {
	return g.CreateEdgeWeight(from, to, 0)
}

// Creates an edge in the graph with a corresponding weight.
// This function will return an error if either of the nodes
// do not belong in the graph.
// Calling CreateEdgeWeight multiple times on the same nodes will not
// make multiple edges; this function will update
// the weight on the node to a new value.
//
// Runs in O(E) time, where E is the number of edges coming out
// of the from node (and to node if the graph is undirected).
// If the graph is sparse, which it should be if you are using
// an adjacency list graph structure like this package provides,
// then connecting two nodes will be near O(1) time. For the exact
// specifics on the performance of this function, click on it and
// read the function.
func (g *Graph) CreateEdgeWeight(from, to Node, weight int) error {
	if from.node == nil || from.node.index >= len(g.nodes) || g.nodes[from.node.index] != from.node {
		return errors.New("First node in CreateEdge call does not belong to this graph")
	}
	if to.node == nil || to.node.index >= len(g.nodes) || g.nodes[to.node.index] != to.node {
		return errors.New("Second node in CreateEdge call does not belong to this graph")
	}
	for edgeI, edge := range from.node.edges { // check if edge already exists
		if edge.end == to.node {
			from.node.edges[edgeI].weight = weight
			// fix reversed edge's weight as well
			for rEdgeI, rEdge := range to.node.reversedEdges { // TODO: add test for this
				if rEdge.end == from.node { // (priority low, reversed weight
					to.node.reversedEdges[rEdgeI].weight = weight //  should not be used anyway)
				}
			}
			return nil
		}
	}
	newEdge := edge{weight: weight, end: to.node}
	// reversed edges only used in directed graph algorithms
	reversedEdge := edge{weight: weight, end: from.node}
	from.node.edges = append(from.node.edges, newEdge)
	to.node.reversedEdges = append(to.node.reversedEdges, reversedEdge)
	if g.kind == Undirected && to != from {
		to.node.edges = append(to.node.edges, reversedEdge)
	}
	return nil
}

// Removes any edges between the nodes. Runs in O(E) time.
func (g *Graph) RemoveEdge(from, to Node) {
	fromEdges := from.node.edges
	toEdges := to.node.edges
	toREdges := to.node.reversedEdges
	for i, edge := range fromEdges {
		if edge.end == to.node {
			fromEdges[i], fromEdges[len(fromEdges)-1] = fromEdges[len(fromEdges)-1], fromEdges[i]
			fromEdges[len(fromEdges)-1].end = nil // for garbage collection
			fromEdges = fromEdges[:len(fromEdges)-1]
			from.node.edges = fromEdges
			break
		}
	}
	for i, edge := range toREdges {
		if edge.end == from.node {
			toREdges[i], toREdges[len(toREdges)-1] = toREdges[len(toREdges)-1], toREdges[i]
			toREdges[len(toREdges)-1].end = nil
			toREdges = toREdges[:len(toREdges)-1]
			to.node.reversedEdges = toREdges
			break
		}
	}
	if g.kind == Undirected && from.node != to.node {
		for i, edge := range toEdges {
			if edge.end == from.node {
				toEdges[i], toEdges[len(toEdges)-1] = toEdges[len(toEdges)-1], toEdges[i]
				toEdges[len(toEdges)-1].end = nil
				toEdges = toEdges[:len(toEdges)-1]
				to.node.edges = toEdges
				break
			}
		}
	}
}

// Returns a slice of nodes that are reachable from the given node in a graph.
func (g *Graph) Neighbors(n Node) []Node {
	neighbors := make([]Node, 0, len(n.node.edges))
	if g.nodes[n.node.index] == n.node {
		for _, edge := range n.node.edges {
			neighbors = append(neighbors, edge.end.container)
		}
	}
	return neighbors
}
