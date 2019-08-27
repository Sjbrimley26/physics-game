package main

// Shape represents a shape. c'mon.
type Shape interface {
	Vertices() []Point
}

// Node actually represents a parent node, a node containing other nodes.
type Node struct {
	box      Rectangle
	children []Shape
}

// IndexNode represents a node on the RTree containing 1 or more LeafNodes.
type IndexNode struct {
	box      Rectangle
	children []interface{}
}

// LeafNode repesents a node on the RTree containing a single object.
type LeafNode struct {
	box Rectangle
	obj Shape
}

// RTree repesents a grouping of objects by location.
type RTree struct {
	box      Rectangle
	children []interface{}
}

func createRTree(center Point, length float64, width float64) RTree {
	return RTree{Rectangle{center, length, width, 0}, make([]interface{}, 3)}
}

func (n *Node) insertShape(s Shape) {
	mbb := MBB(shape.Vertices())

	if len(n.children) < 3 {
		leaf := LeafNode{mbb, s}
		append(n.children, leaf)
		return
	}

	var eligibleParents []interface{}

	for _, child := range n.children {
		if t, is := child.(LeafNode); is == true {
			append(eligibleParents, child)
		}
	}

	if len(eligibleParents) == 0 {
		for _, child := range n.children {
			if len(child.children) < 3 {
				append(eligibleParents, child)
			}
		}
	}

	if len(eligibleParents) == 0 {
		eligibleParents = n.children
	}

	type BoxNode struct {
		child interface{}
		box   Rectangle
	}

	var potentialBoxes []BoxNode
	for _, child := range eligibleParents {
		box := child.box
		pnts := box.Vertices()
		append(pnts, mbb.Vertices())
		expanded := MBB(pnts)
		append(potentialBoxes, BoxNode{child, expanded})
	}

	var smallestBox BoxNode
	for i, n := range potentialBoxes {
		if i == 0 {
			smallestBox = n
			continue
		}
		if n.box.Perimeter() < smallestBox.Perimeter() {
			smallestBox = n
		}
	}

	var childIndex int
	for i, child := range n.children {
		if child == smallestBox.child {
			childIndex = i
			return
		}
		childIndex = -1
	}

	if t, is := n.children[childIndex].(IndexNode); is == true {
		n.children[childIndex].insertShape(s)
		n.children[childIndex].box = smallestBox.box
		return
	}

	child := n.children[childIndex]
	node := IndexNode{smallestBox.box, make([]interface{}, 0)}
}
