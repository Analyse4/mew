// Directed Graph, No Generic
package graph

type Node struct {
	Value    int
	Children []*Node
}

type Graph struct {
	v     int
	e     int
	Nodes []*Node
}

func New(v int) *Graph {
	g := &Graph{v, 0, make([]*Node, 0)}
	for i := 0; i < v; i++ {
		g.Nodes = append(g.Nodes, &Node{0, make([]*Node, 0)})
	}
	return g
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v, w int) {
	g.Nodes[v].Children = append(g.Nodes[v].Children, g.Nodes[w])
	g.e++
}

func (g *Graph) Adj(v int) []*Node {
	return g.Nodes[v].Children
}
