package tree

import "github.com/davecgh/go-spew/spew"

type Connector struct {
	Prev *Node
	V    int
	Next *Node
}

type Node struct {
	V          int
	Connectors []*Connector
}

type Tree struct {
	Root *Node
	Size int
}

func NewNode(v int) *Node {
	n := &Node{
		V: v,
	}

	var connectors []*Connector
	for i := 1; i <= v; i++ {
		connectors = append(connectors, &Connector{
			Prev: n,
			V:    i,
		})
	}

	n.Connectors = connectors

	return n
}

func NewLeaf() *Node {
	n := &Node{
		V: -1,
	}

	return n
}

func buildTree(nums []int, n *Node, idx int) {
	compIdx := len(nums)
	finIdx := compIdx - 1
	for _, conn := range n.Connectors {
		i := idx + conn.V

		var n *Node
		if i > finIdx {
			n = NewLeaf()
		} else {
			v := nums[i]
			n = NewNode(v)
		}

		conn.Next = n
		buildTree(nums, n, i)
	}

}

func NewTree(nums []int) *Tree {
	root := NewNode(nums[0])
	t := &Tree{
		Root: root,
	}
	buildTree(nums, t.Root, 0)
	return t
}

type result struct {
	V      *bool
	Values []int
	Jumps  []int
}

func (r *result) setValue(v bool) *result {
	r.V = &v

	return r
}

func (t *Tree) roundTrip(fin *[]*result, n *Node, r *result) {

	if n.V == -1 {
		r.setValue(true)
		*fin = append(*fin, r)
		//spew.Dump("fin: ", fin)
	}

	if n.V == 0 {
		r.setValue(false)
		*fin = append(*fin, r)
		//spew.Dump("fin: ", fin)
	}

	if len(n.Connectors) == 4 {
		spew.Dump("node:", n)
		spew.Dump("connectors:", n.Connectors)
	}

	for _, conn := range n.Connectors {
		nextNode := conn.Next
		res := &result{
			Values: append(r.Values, nextNode.V),
			Jumps:  append(r.Jumps, conn.V),
		}

		t.roundTrip(fin, nextNode, res)
	}
}

func (t *Tree) RoundTrip() *[]*result {
	n := t.Root
	res := &result{
		V:      nil,
		Values: []int{n.V},
		Jumps:  []int{},
	}

	var fin = []*result{res}

	t.roundTrip(&fin, n, res)
	spew.Dump("finRes: ", fin)

	return &fin
}
