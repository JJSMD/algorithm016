/*
589. N叉树的前序遍历
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
*/

package main

func main() {

}

//遍历
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			queue = append(queue, node.Children...)
		}
		res = append(res, temp)
	}
	return res
}
