/*
33. 搜索旋转排序数组
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
*/
//2020/10/04

package main

func main() {

	test:=[]int{4,5,6,7,0,1,2}
	search(test,0)

}
//
//和一般二分的区别，在于数组被分成了两段单调增的，所以需要在切分时判断是在那个一个递增数组中

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[lo] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
