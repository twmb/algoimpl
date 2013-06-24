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
	// Value can be used to store information on the caller side.
	// Its use is optional. See the Topological Sort example for
	// a reason on why to use this pointer.
	// The reason it is a pointer is so that graph function calls
	// can test for equality on Nodes. The pointer wont change,
	// the value it points to will. If the pointer is explicitly changed,
	// graph functions that use Nodes will cease to work.
	Value *interface{}
}

// An edge connects two Nodes in a graph. To modify the weight, use
// the ConnectWeight function. Any local modifications will
// not be seen in the graph.
//
// In an undirected graph, the start of an edge and end of an edge
// is represented once in the graph: if you connect A to B
// and use the Remove function to remove B, the returned edge will
// have a Start of A and an End of B.
type Edge struct {
	Weight int
	Start  Node
	End    Node
	Kind   GraphType
}

// Prints in the following format:
//     {5 2->3}
// Where 5 is the weight and 2->3 is an edge.
func (e Edge) String() string {
	rVal := fmt.Sprintf("{%v", e.Weight)
	rVal += " " + e.Start.node.String() + "->"
	rVal += e.End.node.String() + "}"
	return rVal
}

// Creates and returns an empty graph. This function must be called before nodes can be connected.
// If kind is Directed, returns a directed graph.
// If kind is Undirected, this function will return an undirected graph.
// If kind is anything else, this function will return an undirected graph by default.
func New(kind GraphType) *Graph {
	switch kind {
	case Directed:
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge), kind: Directed}
	default:
		return &Graph{nodes: []*node{}, edges: make(map[*node][]Edge)}
	}
}

// Creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	newNode := &node{graphIndex: len(g.nodes)}
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
	for _, node := range g.nodes {
		edges := g.edges[node]
		// O(E)
		swapIndex := -1 // index edge to remove is at: swap this with element at end of slice
		for edgeI, edge := range edges {
			if edge.Start == *remove || edge.End == *remove {
				nodeExists = true
				edge.Start.node = nil
				edge.End.node = nil
				swapIndex = edgeI
			}
		}
		if swapIndex > -1 {
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

// Creates an edge in the graph. This function will return an error
// if either of the nodes do not belong in the graph.
//
// Calling connect multiple times on the same nodes will not
// make multiple edges. Because of this existence check, connecting nodes runs in O(N) time.
func (g *Graph) Connect(from, to Node) error {
	return g.ConnectWeight(from, to, 0)
}

// This function calls Connect with the from and to nodes.
// If Connect returns nil, this function will return nil.
// Otherwise, this function sets the weight on an edge
// and then returns a pointer to that edge.
//
// Calling ConnectWeight multiple times on the same nodes will not
// make multiple edges; this function will just update
// the weight on the node to a new value.
func (g *Graph) ConnectWeight(from, to Node, weight int) error {
	if from.node.graphIndex >= len(g.nodes) || g.nodes[from.node.graphIndex] != from.node {
		return errors.New("First node in Connect call does not belong to this graph")
	}
	if to.node.graphIndex >= len(g.nodes) || g.nodes[to.node.graphIndex] != to.node {
		return errors.New("Second node in Connect call does not belong to this graph")
	}
	for edgeI, edge := range g.edges[from.node] { // check if edge already exists
		if edge.End == to || edge.Start == to && edge.End == from {
			g.edges[from.node][edgeI].Weight = weight
			return nil
		}
	}
	newEdge := Edge{Weight: weight, Start: from, End: to}
	if g.kind == Directed {
		newEdge.Kind = Directed
	}
	g.edges[from.node] = append(g.edges[from.node], newEdge)
	if g.kind == Undirected && to != from { // for book keeping
		g.edges[to.node] = append(g.edges[to.node], newEdge)
	}
	return nil
}

// Removes any edges between the nodes. Runs in O(E) time.
func (g *Graph) Unconnect(from, to Node) {
	fromEdges := g.edges[from.node]
	toEdges := g.edges[to.node]
	for i, edge := range fromEdges {
		if edge.Start == to && edge.End == from || edge.End == to && edge.Start == from {
			fromEdges[i], fromEdges[len(fromEdges)-1] = fromEdges[len(fromEdges)-1], fromEdges[i]
			fromEdges = fromEdges[:len(fromEdges)-1]
			g.edges[from.node] = fromEdges
			break
		}
	}
	if g.kind == Undirected {
		for i, edge := range toEdges {
			if edge.Start == from && edge.End == to || edge.End == from && edge.Start == to {
				toEdges[i], toEdges[len(toEdges)-1] = toEdges[len(toEdges)-1], toEdges[i]
				toEdges = toEdges[:len(toEdges)-1]
				g.edges[to.node] = toEdges
				break
			}
		}
	}
}
