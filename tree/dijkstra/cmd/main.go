package main

import (
	"fmt"

	g "github.com/spv-dev/algorithms/tree/dijkstra/internal/graph"
	d "github.com/spv-dev/algorithms/tree/dijkstra/pkg/dijkstra_v1"
)

func main() {
	graph := g.NewGraph()

	graph.AddEdge("book", "plast", 5)
	graph.AddEdge("book", "poster", 0)
	graph.AddEdge("plast", "baraban", 20)
	graph.AddEdge("plast", "guitar", 15)
	graph.AddEdge("poster", "guitar", 30)
	graph.AddEdge("guitar", "piano", 20)
	graph.AddEdge("baraban", "piano", 10)
	graph.AddEdge("piano", "", 0)

	costs := d.Costs(graph, "book")
	fmt.Println(costs)
	fmt.Println(costs["piano"])
}
