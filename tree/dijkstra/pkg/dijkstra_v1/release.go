package dijkstra

import (
	"math"

	g "github.com/spv-dev/algorithms/tree/dijkstra/internal/graph"
)

func Costs(graph *g.Graph, start string) map[string]int {
	costs := make(map[string]int)
	visited := make(map[string]struct{})

	for node := range graph.Vertices {
		costs[node] = math.MaxInt32
	}

	costs[start] = 0

	for range graph.Vertices {
		min := minCost(costs, visited)
		visited[min] = struct{}{}

		for v, price := range graph.Vertices[min] {
			_, ok := visited[v]
			if !ok && costs[min]+price < costs[v] {
				costs[v] = costs[min] + price
			}
		}
	}

	return costs
}

func minCost(neighbours map[string]int, visited map[string]struct{}) string {
	min := math.MaxInt32
	minNode := ""

	for node, dist := range neighbours {
		_, ok := visited[node]
		if !ok && dist < min {
			min = dist
			minNode = node
		}
	}

	return minNode
}
