package graph

const (
	unseen = iota
	visited
)

func dfs(node *Node, finishList *[]*Node) {
	node.state = visited
	for _, adjacentNode := range node.adjacent {
		if adjacentNode.state == unseen {
			adjacentNode.parent = node
			dfs(adjacentNode, finishList)
		}
	}
	*finishList = append(*finishList, node)
}

// Topologically sorts a directed, acyclic graph.
// If the graph is cyclic, the sort order will change
// based on which node the sort starts on.
func TopologicalSort(g *Graph) []*Node {
	sorted := make([]*Node, 0)
	// sort preorder (first jacket, then shirt)
	for _, node := range g.nodes {
		if node.state == unseen {
			dfs(node, &sorted)
		}
	}
	// now make post order for correct sort (jacket follows shirt)
	length := len(sorted)
	for i := 0; i < length/2; i++ {
		sorted[i], sorted[length-i-1] = sorted[length-i-1], sorted[i]
	}
	return sorted
}
