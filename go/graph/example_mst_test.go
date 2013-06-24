package graph_test

import (
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
)

func ExampleMinimumSpanningTree() {
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
	edges := make([]graph.Edge, 0)
	edges = append(edges, *g.Connect(nodes['a'], nodes['b']))
	edges = append(edges, *g.Connect(nodes['a'], nodes['h']))
	edges = append(edges, *g.Connect(nodes['b'], nodes['c']))
	edges = append(edges, *g.Connect(nodes['b'], nodes['h']))
	edges = append(edges, *g.Connect(nodes['c'], nodes['d']))
	edges = append(edges, *g.Connect(nodes['c'], nodes['f']))
	edges = append(edges, *g.Connect(nodes['c'], nodes['i']))
	edges = append(edges, *g.Connect(nodes['d'], nodes['e']))
	edges = append(edges, *g.Connect(nodes['d'], nodes['f']))
	edges = append(edges, *g.Connect(nodes['e'], nodes['f']))
	edges = append(edges, *g.Connect(nodes['f'], nodes['g']))
	edges = append(edges, *g.Connect(nodes['g'], nodes['h']))
	edges = append(edges, *g.Connect(nodes['g'], nodes['i']))
	edges = append(edges, *g.Connect(nodes['h'], nodes['i']))
	*edges[0].Weight = 4
	*edges[1].Weight = 8
	*edges[2].Weight = 8
	*edges[3].Weight = 11
	*edges[4].Weight = 7
	*edges[5].Weight = 4
	*edges[6].Weight = 2
	*edges[7].Weight = 9
	*edges[8].Weight = 14
	*edges[9].Weight = 10
	*edges[10].Weight = 2
	*edges[11].Weight = 1
	*edges[12].Weight = 6
	*edges[13].Weight = 7
	mst := g.MinimumSpanningTree()
	weightSum := 0
	for i := range mst {
		weightSum += *mst[i].Weight
	}
	fmt.Println(weightSum)
	// Output: 37
}
