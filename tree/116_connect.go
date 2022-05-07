package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	link(root.Left, root.Right)
	return root
}

func link(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.Next = right

	link(left.Left, left.Right)
	link(right.Left, right.Right)
	link(left.Right, right.Left)
}
