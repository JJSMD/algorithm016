/*
590. N叉树的后序遍历
https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
*/

package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	for _, cur := range root.Children {
		res = append(res, postorder(cur)...)
	}
	res = append(res, root.Val)
	return res
}

//迭代
func postorder1(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack, node.Children...)
	}

	l := len(res) - 1
	for i := 0; i < l/2+1; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}

	return res
}
