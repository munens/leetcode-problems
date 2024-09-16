package tree

type Connector struct {
	Prev  *Node
	JumpV int
	Next  *Node
}

type Node struct {
	V          int
	Connectors []*Connector
	IsLeaf     bool
}

type Tree struct {
	Root      *Node
	NodeCount int
	HasLeaf   bool
}

func (t *Tree) NewNode(v int, isLeaf bool) *Node {

	if isLeaf {
		if t.HasLeaf != true {
			t.HasLeaf = true
		}
	}

	n := &Node{
		V:      v,
		IsLeaf: isLeaf,
	}

	var connectors []*Connector
	for i := 1; i <= v; i++ {
		connectors = append(connectors, &Connector{
			Prev:  n,
			JumpV: i,
		})
	}

	n.Connectors = connectors

	t.NodeCount += 1
	return n
}

func (t *Tree) NewLeaf() *Node {
	n := &Node{
		V: -1,
	}
	t.NodeCount += 1

	return n
}

func (t *Tree) build(nums []int, n *Node, idx int) {
	finIdx := len(nums) - 1
	for _, conn := range n.Connectors {
		i := idx + conn.JumpV

		var n *Node
		if i > finIdx {
			n = t.NewNode(-1, true)
		} else {
			v := nums[i]
			n = t.NewNode(v, i == finIdx)
		}

		conn.Next = n

		t.build(nums, n, i)
	}
}

func (t *Tree) canBuildToLeaf(nums []int, n *Node, idx int) bool {
	finIdx := len(nums) - 1
	for _, conn := range n.Connectors {
		i := idx + conn.JumpV

		var n *Node
		if i > finIdx {
			n = t.NewNode(-1, true)
			return true
		} else {
			v := nums[i]
			n = t.NewNode(v, i == finIdx)

			if i == finIdx {
				return true
			}
		}

		conn.Next = n

		return t.canBuildToLeaf(nums, n, i)
	}

	return false
}

func NewTree(nums []int) *Tree {
	t := &Tree{}
	root := t.NewNode(nums[0], len(nums) == 1 && nums[0] == 0)
	t.Root = root
	t.build(nums, t.Root, 0)
	return t
}

func CanBuildToTree(nums []int) bool {
	t := &Tree{}
	root := t.NewNode(nums[0], len(nums) == 1 && nums[0] == 0)
	t.Root = root
	return t.canBuildToLeaf(nums, t.Root, 0)
}

type result struct {
	V      bool
	Values []int
	Jumps  []int
}

func (t *Tree) roundTrip(fin *[]*result, n *Node, r *result) {

	if n.IsLeaf {
		r.V = true
		f := *r
		*fin = append(*fin, &f)
		return
	}

	for _, conn := range n.Connectors {
		nextNode := conn.Next

		// Make a copy of r.Values
		values := make([]int, len(r.Values))
		copy(values, r.Values)
		values = append(values, nextNode.V)

		// Make a copy of r.Jumps
		jumps := make([]int, len(r.Jumps))
		copy(jumps, r.Jumps)
		jumps = append(jumps, conn.JumpV)

		res := &result{
			Values: values,
			Jumps:  jumps,
		}

		t.roundTrip(fin, nextNode, res)
	}
}

func (t *Tree) RoundTrip() *[]*result {
	n := t.Root
	res := &result{
		V:      false,
		Values: []int{n.V},
		Jumps:  []int{},
	}

	var fin []*result

	t.roundTrip(&fin, n, res)

	return &fin
}

func CanJump(nums []int) bool {
	return CanBuildToTree(nums)
}
