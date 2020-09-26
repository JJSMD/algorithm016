/*
22. 括号生成
https://leetcode-cn.com/problems/generate-parentheses/
*/
//2020/9/23

package main

func main() {

}

//递归 深度优先
func generateParenthesis(n int) []string {
	list := new([]string)
	//backtrack(list, 0, 0, n, "")
	backtrack1(list, n, n, "")
	return *list
}

func backtrack(list *[]string, left int, right int, n int, input string) {
	if left > n || left < right {
		return
	}

	if left+right == 2*n {
		*list = append(*list, input)
		return
	}

	backtrack(list, left+1, right, n, input+"(")
	backtrack(list, left, right+1, n, input+")")
}

func backtrack1(list *[]string, left int, right int, input string) {
	if right == 0 {
		*list = append(*list, input)
		return
	}

	if left > 0 {
		backtrack1(list, left-1, right, input+"(")
	}

	if right > left {
		backtrack1(list, left, right-1, input+")")
	}
}
