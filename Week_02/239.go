/*
239. 滑动窗口最大值
https://leetcode-cn.com/problems/sliding-window-maximum/
*/
//2020/09/20

package main

func main() {
	input := []int{1,3,-1,-3,5,3,6,7}
	maxSlidingWindow(input,3)
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := []int{}

	cleanDeque := func(i int) {
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
	}

	for i := 0; i < k; i++ {
		cleanDeque(i)
		deque = append(deque, i)
	}
	res = append(res, nums[deque[0]])

	for i := k; i < len(nums); i++ {
		cleanDeque(i)
		deque = append(deque,i)
		res = append(res,nums[deque[0]])
	}
	return res
}
