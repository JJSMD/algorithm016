##### 树常用算法

###### 二叉树的前中后序

1. 深度优先遍历
   - 前序遍历，根-左-右
   - 中序遍历，左-根-右
   - 后序遍历，左-右-根

2. 深度优先遍历
   - 层序遍历，逐层遍历并返回

###### 主要遍历方法有，递归法、迭代法：

​	递归法比较好理解，但效率最低一般不会真正使用，代码通用只是前中后序的递归顺序不一样,递归法基本就是一套代码下来。

```go
func traversal(root *TreeNode) []int {
	if root == nil {
            return nil
        }
	res := []int{}
	//前序
	res = append(res, root.Val)
	res = append(res, traversal(root.Left)...)
	res = append(res, traversal(root.Right)...)
	
	//中序
	res = append(res, traversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, traversal(root.Right)...)
	
	//后续
	res = append(res, traversal(root.Left)...)
	res = append(res, traversal(root.Right)...)
	res = append(res, root.Val)	

	return res
}
```

​	迭代法，在实现深度优先的前中后序遍历时，主要借助栈来模拟递归。但前中后序代码并不能形成统一写法，主要是因为无法同时解决处理过程和访问过程不一致的情况。所以主要是掌握思路，实现起来并不困难。



###### N叉树的前序、后序遍历（思路和方法基本可以完全参考二叉树，需要注意对子树入栈的区别处理）



###### 