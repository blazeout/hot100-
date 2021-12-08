package _1_树的最大距离

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 需要遍历每个节点左边最大的节点数和右边最大的节点数之和即可, 同时更新maxLength的值
func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDeep := dfs(node.Left)
		rightDeep := dfs(node.Right)
		if leftDeep+rightDeep > maxDiameter {
			maxDiameter = leftDeep + rightDeep
		}
		return max(leftDeep, rightDeep) + 1
	}
	dfs(root)
	return maxDiameter
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
