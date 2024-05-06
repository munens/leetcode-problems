package tree

type Connector struct {
	Start *Node
	V     int
	End   *Node
}

type Node struct {
	V          int
	Connectors []*Connector
}

type Tree struct {
	Root *Node
}

func NewNode(v int, addConnectors bool) *Node {
	n := &Node{
		V: v,
	}

	var connectors []*Connector
	if addConnectors {
		for i := 1; i <= v; i++ {
			connectors = append(connectors, &Connector{
				Start: n,
				V:     i,
			})
		}
	}

	n.Connectors = connectors

	return n
}

func NewTree(rootV int) *Tree {
	root := NewNode(rootV, true)
	return &Tree{
		Root: root,
	}
}
