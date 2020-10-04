#### 搜索

深度优先搜索模板：

```go
func DFS(root *TreeNode) []int {
	visited := make(map[TreeNode]bool)
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []TreeNode{*root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[node]; ok {
			res = append(res, node.Val)
			continue
		}
		visited[node] = true

		if node.Left != nil {
			stack = append(stack, *node.Left)
		}
		if node.Right != nil {
			stack = append(stack, *node.Right)
		}
	}
	return res
}
```

广度优先搜索模板：

```go
func returnBFS(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []TreeNode{*root}
	for len(queue) > 0 {
		currLen := len(queue)
		temp := []int{}
		for i := 0; i < currLen; i++ {
			node := queue[0]
			queue = queue[1 : len(queue)-1]
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

```



#### 贪心算法

贪心算法是在每一步的选择中，选择当前状态下最优解，从而导致结果是全局的最优结果。贪心算法没有回退功能。

运用中，一旦一个问题可以通过贪心算法来解决，那么贪心一般就是解决这个问题的最好办法。由于贪心算法的高效性以及其所求解的答案比较接近最优结果，所以贪心一般可以作为辅助算法，或者直接解决一些要求结果不是特别精确的问题。

贪心适用的一些问题（最优解问题）：

- 求K个数最小值问题
- 求图中最小生成树问题
- 求霍夫曼编码

#### 二分查找

二分查找是比较高效的查找方法，但是有一些先决条件：

- 必须是顺序存储结构

- 必须单调递增或单调递减

- 必须存在上下界

  

二分查找模板：

```go
func bSearch(nums []int, target int) {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			// return result
			return
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}
```

