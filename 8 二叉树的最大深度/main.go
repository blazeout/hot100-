package __二叉树的最大深度

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树一般是DFS或者BFS DFS递归实现, BFS队列实现
func maxDepth1(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	// 树的深度 = 1 + max(l,r)
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(dfs(root.Left), dfs(root.Right))
	}
	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// BFS队列实现
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		count++
	}
	return count
}
