/*
589. N叉树的前序遍历
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
*/

package main

func main() {

}

//迭代
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res, root.Val)
	for _, cur := range root.Children {
		res = append(res, preorder(cur)...)
	}

	return res
}

//迭代
func preorder1(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		l := len(node.Children)
		for i := l - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}

	}

	return res
}
