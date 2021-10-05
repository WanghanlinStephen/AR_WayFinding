package structure

import "fmt"

type Label struct {
	Id string
	NameEnglish string
	NameChinese string
	NameChineseTradition string
	Latitude string
	Longitude string
}

type Edge struct {
	Neighbors map[Label]struct{}
	EstimatedTime int
}

type Graph struct {
	Nodes map[Label]struct{}
	Edges map[Label]Edge
}

// NewGraph: Create graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[Label]struct{}),
		Edges: make(map[Label]Edge),
	}
}

// AddNode: Add node id to graph, return true if added (Label's are unique).
func (g *Graph) AddNode(id Label) bool {
	if _, ok := g.Nodes[id]; ok {
		return false
	}
	g.Nodes[id] = struct{}{}
	return true
}

// AddNodes: Add node ids to graph, return true if added (Label's are unique).
func (g *Graph) AddNodes(ids []Label) bool {
	for _,value := range ids {
		if ok := g.AddNode(value);!ok{
			//fixme:记录日志
		}
	}
	return true
}

// AddEdge: Add an edge from u to v.
func (g *Graph) AddEdge(u, v Label) {
	if _, ok := g.Nodes[u]; !ok {
		g.AddNode(u)
	}
	if _, ok := g.Nodes[v]; !ok {
		g.AddNode(v)
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = Edge{make(map[Label]struct{}),0}
	}
	g.Edges[u].Neighbors[v]=struct{}{}
	// For undirected graph add edge from v to u.
	if _, ok := g.Edges[v]; !ok {
		g.Edges[v] = Edge{make(map[Label]struct{}),0}
	}
	g.Edges[v].Neighbors[u]=struct{}{}
}


// Print Adjacent eldges
func (g *Graph) adjacentEdgesExample(start Label) {
	fmt.Printf("Printing all edges adjacent to %s\n", start.Id)
	for v := range g.Edges[start].Neighbors {
		// Edge exists from u to v.
		fmt.Printf("Edge: %s -> %s\n", start, v.Id)
	}

	fmt.Println("Printing all edges.")
	for u, m := range g.Edges {
		for v := range m.Neighbors {
			// Edge exists from u to v.
			fmt.Printf("Edge: %s -> %s\n", u.Id, v.Id)
		}
	}
}