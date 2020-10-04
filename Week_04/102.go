/*
46. 全排列
https://leetcode-cn.com/problems/permutations/
*/
//2020/10/01

package main

func main() {

}

//遍历
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []TreeNode{}
	queue = append(queue, *root)

	for len(queue) > 0 {
		temp := []int{}
		currLen := len(queue)
		for i := 0; i < currLen; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}
			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
