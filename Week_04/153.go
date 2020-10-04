/*
153. 寻找旋转排序数组中的最小值
https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
*/
//2020/10/04

package main

func main() {

}

func findMin(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	l, r := 0, len(nums)-1

	for l < r {
		//mid := (l + r) >> 1
		//优化 实际上min更靠近左边界
		mid := l + (r-l)>>2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

//直接遍历
func findMin1(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}
	return nums[0]
}
