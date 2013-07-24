package graph_test

import (
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
)

func ExampleGraph_MinimumSpanningTree() {
	g := graph.New(graph.Undirected)
	nodes := make(map[rune]graph.Node, 0)
	nodes['a'] = g.MakeNode()
	nodes['b'] = g.MakeNode()
	nodes['c'] = g.MakeNode()
	nodes['d'] = g.MakeNode()
	nodes['e'] = g.MakeNode()
	nodes['f'] = g.MakeNode()
	nodes['g'] = g.MakeNode()
	nodes['h'] = g.MakeNode()
	nodes['i'] = g.MakeNode()
	g.CreateEdgeWeight(nodes['a'], nodes['b'], 4)
	g.CreateEdgeWeight(nodes['a'], nodes['h'], 8)
	g.CreateEdgeWeight(nodes['b'], nodes['c'], 8)
	g.CreateEdgeWeight(nodes['b'], nodes['h'], 11)
	g.CreateEdgeWeight(nodes['c'], nodes['d'], 7)
	g.CreateEdgeWeight(nodes['c'], nodes['f'], 4)
	g.CreateEdgeWeight(nodes['c'], nodes['i'], 2)
	g.CreateEdgeWeight(nodes['d'], nodes['e'], 9)
	g.CreateEdgeWeight(nodes['d'], nodes['f'], 14)
	g.CreateEdgeWeight(nodes['e'], nodes['f'], 10)
	g.CreateEdgeWeight(nodes['f'], nodes['g'], 2)
	g.CreateEdgeWeight(nodes['g'], nodes['h'], 1)
	g.CreateEdgeWeight(nodes['g'], nodes['i'], 6)
	g.CreateEdgeWeight(nodes['h'], nodes['i'], 7)
	mst := g.MinimumSpanningTree()
	weightSum := 0
	for i := range mst {
		weightSum += mst[i].Weight
	}
	fmt.Println(weightSum)
	// Output: 37
}
