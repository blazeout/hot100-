package __二叉树的中序遍历

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归是操作系统/虚拟机隐式帮我们维护一个栈
/*func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    nums := []int{}
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode){
        if root == nil {
            return
        }
        inorder(root.Left)
        nums = append(nums, root.Val)
        inorder(root.Right)
    }
    inorder(root)
    return nums
}*/

// 迭代需要我们显示的维护一个栈
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 相当于左遍历
			stack = append(stack, root)
			root = root.Left
		}
		// 中间的操作
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 右遍历
		root = root.Right
	}
	return res
}
