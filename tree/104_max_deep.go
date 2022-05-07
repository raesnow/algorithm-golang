package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	depth := 0
	for len(queue) != 0 {
		newQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			if node != nil {
				newQueue = append(newQueue, node.Left)
				newQueue = append(newQueue, node.Right)
			}
		}
		if len(newQueue) != 0 {
			depth++
		}
		queue = newQueue
	}
	return depth
}
