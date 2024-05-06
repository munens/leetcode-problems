package main

type Node struct {
	v        int
	prev     *Node
	children []*Node
}

func (n *Node) AddChild(v int) {
	child := &Node{
		v:        v,
		prev:     n,
		children: []*Node{},
	}
	n.children = append(n.children, child)
}

type Tree struct {
	root *Node
}

func (t *Tree) Create(rootV int) *Tree {
	root := &Node{
		v:        rootV,
		prev:     nil,
		children: []*Node{},
	}
	t.root = root

	return t
}

func (t *Tree) IsEmpty() bool {
	return t.root == nil
}
