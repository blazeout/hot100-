package _5_反转二叉树



// TreeNode Definition for a binary tree node.
type TreeNode struct {
   Val int
  Left *TreeNode
    Right *TreeNode
}


// dfs 深度优先遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode)  {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)
	return root
}
