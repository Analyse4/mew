package tree

type Node struct {
	Value    interface{}
	Children []*Node
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return new(Tree)
}

// TODO
func (t *Tree) Height() int {
	return 0
}
