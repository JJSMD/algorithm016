/*
94. 二叉树的中序遍历
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/

package main

func main() {

}

//递归
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return res
}

//颜色标记法
func inorderTraversal1(root *TreeNode) []int {
	type node struct {
		Color int
		Node  *TreeNode
	}

	stack := make([]node, 0)
	rootNode := node{0, root}

	stack = append(stack, rootNode)

	result := make([]int, 0)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if n.Node == nil {
			continue
		}

		if n.Color == 0 {
			stack = append(stack, node{0, n.Node.Right})
			stack = append(stack, node{1, n.Node})
			stack = append(stack, node{0, n.Node.Left})
		} else {
			result = append(result, n.Node.Val)
		}

	}
	return result
}
