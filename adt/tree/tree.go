package tree

type Node struct {
	value    interface{}
	children []*Node
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return new(Tree)
}
