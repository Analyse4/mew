package graph

import "testing"

func TestGraph(t *testing.T) {
	g := New(5)
	g.AddEdge(1, 4)
	adj := g.Adj(4)
	ra := g.Nodes[1]
	raV := 5
	raE := 1

	for _, v := range adj {
		if v.Value != ra.Value {
			t.Fatalf("wrong value, got: %v, want: %v\n", v.Value, ra.Value)
		}
	}
	if g.V() != raV {
		t.Fatalf("wrong, got: %v, want: %v\n", g.V(), raV)
	}
	if g.E() != raE {
		t.Fatalf("wrong, got: %v, want: %v\n", g.E(), raE)
	}
}
