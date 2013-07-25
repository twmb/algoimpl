package graph

import (
	"container/heap"
	"github.com/twmb/algoimpl/go/graph/lite"
	"math/rand"
	"time"
)

const (
	unseen = iota
	seen
	queued = 1 << 30 // assume that no weight will be > 2^30
)

// O(V + E). It does not matter to traverse back
// on a bidirectional edge, because any vertex dfs is
// recursing on is marked as visited and won't be visited
// again anyway.
func (g *Graph) dfs(node *node, finishList *[]Node) {
	node.state = seen
	for _, edge := range node.edges {
		if edge.end.state == unseen {
			edge.end.parent = node
			g.dfs(edge.end, finishList)
		}
	}
	*finishList = append(*finishList, node.container)
}

func (g *Graph) dfsReversedEdges(node *node, finishList *[]Node) {
	node.state = seen
	for _, edge := range node.reversedEdges {
		if edge.end.state == unseen {
			edge.end.parent = node
			g.dfsReversedEdges(edge.end, finishList)
		}
	}
	*finishList = append(*finishList, node.container)
}

// Topologically sorts a directed acyclic graph.
// If the graph is cyclic, the sort order will change
// based on which node the sort starts on. O(V+E) complexity.
func (g *Graph) TopologicalSort() []Node {
	sorted := make([]Node, 0, len(g.nodes))
	// sort preorder (first jacket, then shirt)
	for _, node := range g.nodes {
		if node.state == unseen {
			g.dfs(node, &sorted)
		}
	}
	// now make post order for correct sort (jacket follows shirt). O(V)
	length := len(sorted)
	for i := 0; i < length/2; i++ {
		sorted[i], sorted[length-i-1] = sorted[length-i-1], sorted[i]
	}
	return sorted
}

// Returns reversed copy of the directed graph g. O(V+E) complexity.
// This function can be used to copy an undirected graph.
func (g *Graph) Reverse() *Graph {
	reversed := New(Directed)
	if g.kind == Undirected {
		reversed = New(Undirected)
	}
	// O(V)
	for _ = range g.nodes {
		reversed.MakeNode()
	}
	// O(V + E)
	for _, node := range g.nodes {
		for _, edge := range node.edges {
			reversed.MakeEdge(reversed.nodes[edge.end.index].container, reversed.nodes[node.index].container)
		}
	}
	return reversed
}

// Returns a slice of strongly connected nodes on a directed graph.
// If passed an undirected graph, returns nil.
// The returned components have reversed, nonexclusive edges.
// For example, if this is passed the graph
//     a->b, c
//     b->a, c
//     c
// will return components
//     [[c->a, b], [b->a], [a->b]]
// where -> represents the edges that the node contains.
// O(V+E) complexity.
func (g *Graph) StronglyConnectedComponents() [][]Node {
	if g.kind == Undirected {
		return nil
	}
	components := make([][]Node, 0)
	finishOrder := g.TopologicalSort()
	for i := range finishOrder {
		finishOrder[i].node.state = unseen
	}
	for _, sink := range finishOrder {
		if g.nodes[sink.node.index].state == unseen {
			component := make([]Node, 0)
			g.dfsReversedEdges(g.nodes[sink.node.index], &component)
			components = append(components, component)
		}
	}
	return components
}

func (g *Graph) RandMinimumCut(iterations int) []Edge {
	// copy a graph into mock nodes and edges
	// so that I can collapse nodes / edges without losing
	// the true spanning edge's start and end
	rand.Seed(time.Now().Unix())
	var minLiteEdges []lite.Edge
	for iteration := 0; iteration < iterations; iteration++ {

		gLite := lite.NewGraph(len(g.nodes))
		for i := range g.nodes {
			for _, edge := range g.nodes[i].edges {
				gLite[i].Edges = append(gLite[i].Edges, lite.Edge{End: edge.end.index, S: g.nodes[i], E: edge.end})
			}
		}

		for len(gLite) > 2 {
			randNodeI := rand.Intn(len(gLite))
			randNode := gLite[randNodeI]

			randEdgeI := rand.Intn(len(randNode.Edges))
			randEdge := randNode.Edges[randEdgeI]

			removeNodeI := randEdge.End

			finalRandNodeI := randNodeI
			if randNodeI > removeNodeI {
				finalRandNodeI-- // the final index of randNode will be back one
			}

			// every edge on rand node pointing to remove node is removed (will form loop when collapsed)
			shift := 0
			for e := range gLite[randNodeI].Edges {
				if gLite[randNodeI].Edges[e].End == removeNodeI {
					shift++
					continue
				}
				if gLite[randNodeI].Edges[e].End > removeNodeI {
					gLite[randNodeI].Edges[e].End--
				}
				gLite[randNodeI].Edges[e-shift] = gLite[randNodeI].Edges[e]
			}
			gLite[randNodeI].Edges = gLite[randNodeI].Edges[:len(gLite[randNodeI].Edges)-shift]

			// everything on every other node pointing to removeNode should now point to randNode
			for n := range gLite {
				if n == removeNodeI || n == randNodeI {
					continue
				}
				for e := range gLite[n].Edges {
					// if the current edge on the current node points to the removal node
					if gLite[n].Edges[e].End == removeNodeI {
						// point it to the final randNode position
						gLite[n].Edges[e].End = finalRandNodeI
					} else if gLite[n].Edges[e].End > removeNodeI {
						// or, if the edge points to something AFTER the removal node, decrement that edge's end
						gLite[n].Edges[e].End--
					}
				}
			}
			numAppended := 0
			// everything removeNode points to, randNode should now point to
			for e := range gLite[removeNodeI].Edges {
				if gLite[removeNodeI].Edges[e].End != randNodeI {
					if gLite[removeNodeI].Edges[e].End > removeNodeI {
						gLite[removeNodeI].Edges[e].End--
					}
					gLite[randNodeI].Edges = append(gLite[randNodeI].Edges, gLite[removeNodeI].Edges[e])
					numAppended++
				}
			}
			// shift all nodes after the remove node back one
			gLite = append(gLite[:removeNodeI], gLite[removeNodeI+1:]...)
		}
		if iteration == 0 || len(gLite[0].Edges) < len(minLiteEdges) {
			minLiteEdges = make([]lite.Edge, len(gLite[0].Edges))
			copy(minLiteEdges, gLite[0].Edges)
		}
	}

	minCut := make([]Edge, len(minLiteEdges))
	for i := range minLiteEdges {
		start := minLiteEdges[i].S.(*node)
		end := minLiteEdges[i].E.(*node)
		minCut[i] = Edge{Weight: g.edgeWeightBetween(start, end), Start: start.container, End: end.container, Kind: g.kind}
	}
	return minCut
}

// This function will return the edges corresponding to the
// minimum spanning tree in the graph based off of the edge's weight values.
func (g *Graph) MinimumSpanningTree() []Edge {
	// create priority queue for vertices
	nodesBase := nodeSlice(make([]*node, 0))
	nodes := &nodesBase
	heap.Init(nodes)
	for _, node := range g.nodes[1:] {
		node.state = 1<<30 - 1 // bit 31 is for queue testing
		node.state |= queued
		heap.Push(nodes, node)
	}
	g.nodes[0].state = 0
	heap.Push(nodes, g.nodes[0]) // now add root element

	for nodes.Len() > 0 {
		min := heap.Pop(nodes).(*node)
		for _, edge := range min.edges {
			v := edge.end // get the other side of the edge
			if nodes.QueueContains(v) && edge.weight < (v.state & ^queued) {
				v.parent = min
				heap.Remove(nodes, v.data)     // remove it
				v.state = edge.weight | queued // update state, add queue bit
				heap.Push(nodes, v)            // add it back in, WORD, O(lg n) update key time!
			}
		}
	}

	mst := make([]Edge, 0)
	for i := range g.nodes {
		if g.nodes[i].parent != nil {
			mst = append(mst, Edge{Weight: g.edgeWeightBetween(g.nodes[i], g.nodes[i].parent), Start: g.nodes[i].container,
				End: g.nodes[i].parent.container, Kind: g.kind})
		}
	}

	return mst
}

// only called when the graph is guaranteed to have an edge
// between the two nodes
func (g *Graph) edgeWeightBetween(v, u *node) int {
	for _, edge := range u.edges {
		// one of the two is always u
		if edge.end == v {
			return edge.weight
		}
	}
	return 0
}

type nodeSlice []*node

func (n *nodeSlice) Push(x interface{}) {
	p := x.(*node)
	p.data = len(*n) // index into heap
	*n = append(*n, x.(*node))
}
func (n *nodeSlice) Pop() interface{} {
	rNode := (*n)[len(*n)-1]
	rNode.state &= ^queued
	*n = (*n)[0 : len(*n)-1]
	return rNode
}
func (n nodeSlice) Len() int {
	return len(n)
}
func (n nodeSlice) Less(i, j int) bool {
	return n[i].state & ^queued < n[j].state & ^queued
}
func (n nodeSlice) Swap(i, j int) {
	n[j].data, n[i].data = n[i].data, n[j].data // swap data containing indices
	n[j], n[i] = n[i], n[j]
}
func (n nodeSlice) QueueContains(node *node) bool { // extend heap interface
	return node.state&queued > 0
}
