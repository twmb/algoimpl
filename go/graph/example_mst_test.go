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
	g.ConnectWeight(nodes['a'], nodes['b'], 4)
	g.ConnectWeight(nodes['a'], nodes['h'], 8)
	g.ConnectWeight(nodes['b'], nodes['c'], 8)
	g.ConnectWeight(nodes['b'], nodes['h'], 11)
	g.ConnectWeight(nodes['c'], nodes['d'], 7)
	g.ConnectWeight(nodes['c'], nodes['f'], 4)
	g.ConnectWeight(nodes['c'], nodes['i'], 2)
	g.ConnectWeight(nodes['d'], nodes['e'], 9)
	g.ConnectWeight(nodes['d'], nodes['f'], 14)
	g.ConnectWeight(nodes['e'], nodes['f'], 10)
	g.ConnectWeight(nodes['f'], nodes['g'], 2)
	g.ConnectWeight(nodes['g'], nodes['h'], 1)
	g.ConnectWeight(nodes['g'], nodes['i'], 6)
	g.ConnectWeight(nodes['h'], nodes['i'], 7)
	mst := g.MinimumSpanningTree()
	weightSum := 0
	for i := range mst {
		weightSum += mst[i].Weight
	}
	fmt.Println(weightSum)
	// Output: 37
}
