package graph

type Path struct {
	Weight int
	Path   []Edge
}

// DijkstraSearch returns the shortest path from the start node to every other
// node in the graph. All edges must have a positive weight, otherwise this
// function will return nil.
func (g *Graph) DijkstraSearch(start Node) []Path {
	if start.node == nil || g.nodes[start.node.index] != start.node {
		return nil
	}
	paths := make([]Path, len(g.nodes))

	nodesBase := nodeSlice(make([]*node, len(g.nodes)))
	copy(nodesBase, g.nodes)
	for i := range nodesBase {
		nodesBase[i].state = 1<<31 - 1
		nodesBase[i].data = i
	}
	start.node.state = 0 // make it so 'start' sorts to the top of the heap
	nodes := &nodesBase
	nodes.heapInit()

	for len(*nodes) > 0 {
		curNode := nodes.pop()
		for _, edge := range curNode.edges {
			newWeight := curNode.state + edge.weight
			if newWeight < curNode.state { // negative edge length
				return nil
			}
			v := edge.end
			if nodes.heapContains(v) && newWeight < v.state {
				v.parent = curNode
				nodes.update(v.data, newWeight)
			}
		}

		// build path to this node
		if curNode.parent != nil {
			newPath := Path{Weight: curNode.state}
			newPath.Path = make([]Edge, len(paths[curNode.parent.index].Path)+1)
			copy(newPath.Path, paths[curNode.parent.index].Path)
			newPath.Path[len(newPath.Path)-1] = Edge{Weight: curNode.state - curNode.parent.state,
				Start: curNode.parent.container, End: curNode.container}
			paths[curNode.index] = newPath
		} else {
			paths[curNode.index] = Path{Weight: curNode.state, Path: []Edge{}}
		}
	}
	return paths
}

// Context for an AllPathSearch
type allPathSearch struct {
	// Graph
	Graph *Graph
	// Current path
	Path []Edge
	// Set of node values on the path
	NodeValues map[*interface{}]bool
	// Final paths
	Result []Path
}

// AllPathSearch returns all the paths from the start node to another
// node in the graph. The graph can be cyclic but cyclic paths will never be included.
func (g *Graph) AllPathSearch(start, end Node) []Path {
	aps := &allPathSearch{
		Path:       []Edge{},
		NodeValues: make(map[*interface{}]bool),
		Graph:      g,
		Result:     []Path{},
	}
	aps.enumerate(start, end)
	return aps.Result
}

// enumerate recursively evaluates the allPathSearch. Curr is the current visited node, end
// is the target node.
// Implementation is based on: http://introcs.cs.princeton.edu/java/45graph/AllPaths.java.html
func (aps *allPathSearch) enumerate(curr, end Node) {
	// Mark this node as visited
	aps.NodeValues[curr.Value] = true
	// Defer unmarking this as visited when we rewind the stack.
	defer func() {
		delete(aps.NodeValues, curr.Value)
	}()

	// we have found a path to the destination.
	if curr.Value == end.Value {
		path := Path{Weight: 0}
		path.Path = make([]Edge, len(aps.Path))
		copy(path.Path, aps.Path)
		aps.Result = append(aps.Result, path)
		return
	}

	// For each neighbor, continue the path if we haven't already visited that node.
	neighbors := aps.Graph.Neighbors(curr)
	for _, neighbor := range neighbors {
		if _, ok := aps.NodeValues[neighbor.Value]; ok {
			continue
		}
		// Push the edge onto the stack
		aps.Path = append(aps.Path, Edge{Start: curr, End: neighbor})
		aps.enumerate(neighbor, end)
		// Pop the edge off the stack before continuing
		aps.Path = aps.Path[:len(aps.Path)-1]
	}
}
