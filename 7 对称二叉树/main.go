package main

// TreeNode Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	// 1. 当两颗子树同时越过根节点为空时,返回true
	if left == nil && right == nil {
		return true
	}
	// 2. 当一颗左子树或者右子数为空时返回false
	if left == nil || right == nil {
		return false
	}
	// 3. 当遍历的两颗节点值不同时返回false
	if left.Val != right.Val {
		return false
	}
	// 最后返回继续遍历的结果
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}


