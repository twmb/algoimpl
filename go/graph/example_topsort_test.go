package graph_test

import (
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
)

func ExampleGraph_TopologicalSort() {
	g := graph.New(graph.Directed)

	clothes := make(map[string]graph.Node, 0)
	// Make a mapping from strings to a node
	clothes["shirt"] = g.MakeNode()
	clothes["tie"] = g.MakeNode()
	clothes["jacket"] = g.MakeNode()
	clothes["belt"] = g.MakeNode()
	clothes["watch"] = g.MakeNode()
	clothes["undershorts"] = g.MakeNode()
	clothes["pants"] = g.MakeNode()
	clothes["shoes"] = g.MakeNode()
	clothes["socks"] = g.MakeNode()
	// Make references back to the string values
	*clothes["shirt"].Value = "shirt"
	*clothes["tie"].Value = "tie"
	*clothes["jacket"].Value = "jacket"
	*clothes["belt"].Value = "belt"
	*clothes["watch"].Value = "watch"
	*clothes["undershorts"].Value = "undershorts"
	*clothes["pants"].Value = "pants"
	*clothes["shoes"].Value = "shoes"
	*clothes["socks"].Value = "socks"
	// Connect the elements
	g.Connect(clothes["shirt"], clothes["tie"])
	g.Connect(clothes["tie"], clothes["jacket"])
	g.Connect(clothes["shirt"], clothes["belt"])
	g.Connect(clothes["belt"], clothes["jacket"])
	g.Connect(clothes["undershorts"], clothes["pants"])
	g.Connect(clothes["undershorts"], clothes["shoes"])
	g.Connect(clothes["pants"], clothes["belt"])
	g.Connect(clothes["pants"], clothes["shoes"])
	g.Connect(clothes["socks"], clothes["shoes"])
	sorted := g.TopologicalSort()
	for i := range sorted {
		fmt.Println(*sorted[i].Value)
	}
	// Output:
	// socks
	// undershorts
	// pants
	// shoes
	// watch
	// shirt
	// belt
	// tie
	// jacket
}
