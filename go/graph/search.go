package graph

import (
	"container/heap"
)

type Path struct {
	Weight int
	Path   []Edge
}

// DijkstraSearch returns the shortest path to every other node in the graph.
// All edges must have a positive weight, otherwise this function will return
// nil.
func (g *Graph) DijkstraSearch(start Node) []Path {
	paths := make([]Path, len(g.nodes))

	nodesBase := nodeSlice(make([]*node, len(g.nodes)))
	copy(nodesBase, g.nodes)
	for i := range nodesBase {
		nodesBase[i].state = 2<<30 - 1
		nodesBase[i].state |= queued
		nodesBase[i].data = i
	}
	start.node.state = 0 // make it so 'start' sorts to the top of the heap
	start.node.state |= queued
	nodes := &nodesBase
	heap.Init(nodes)

	for nodes.Len() > 0 {
		curNode := heap.Pop(nodes).(*node)
		for _, edge := range curNode.edges {
			newWeight := curNode.state + edge.weight
			v := edge.end
			if nodes.QueueContains(v) && newWeight < (v.state & ^queued) {
				v.parent = curNode
				heap.Remove(nodes, v.data)
				v.state = newWeight | queued
				heap.Push(nodes, v)
			}
		}
		// build path to this node
		if curNode.parent != nil {
			newPath := Path{Weight: curNode.state}
			newPath.Path = make([]Edge, len(paths[curNode.parent.index].Path)+1)
			copy(newPath.Path, paths[curNode.parent.index].Path)
			newPath.Path[len(newPath.Path)-1] = Edge{Weight: curNode.state - curNode.parent.state,
				Start: curNode.parent.container, End: curNode.container, Kind: Directed}
			paths[curNode.index] = newPath
		} else {
			paths[curNode.index] = Path{Weight: curNode.state, Path: []Edge{}}
		}
	}
	return paths
}
