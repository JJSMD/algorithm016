/*
226. 翻转二叉树
https://leetcode-cn.com/problems/invert-binary-tree/description/
*/
//2020/9/23

package main

func main() {

}

//递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//if root.Left == nil && root.Right == nil {
	//	return root
	//}

	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}
