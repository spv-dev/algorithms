package graph

type Graph struct {
	Vertices map[string]map[string]int
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]map[string]int),
	}
}

func (g *Graph) AddEdge(u string, v string, weight int) {
	if g.Vertices[u] == nil {
		g.Vertices[u] = make(map[string]int)
	}
	g.Vertices[u][v] = weight
}
