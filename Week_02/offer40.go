/*
剑指 Offer 40. 最小的k个数
https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
*/

package main

import (
	"sort"
)

func main() {

}

//排序
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}
	return res
}

//思路：
//首先选定第一个元素为基准元素，然后遍历后面的元素，找到比基准元素小的元素并放到基准元素后面，
//最后将基准元素与最后一个比他小的元素交换，如果基准元素的下标正好为k，则直接返回基准元素前半部分数组，
//如果比k大，那么左半部分继续重复上面的过程直到下标k，如果比k，小则右半部分重复上述过程。

func getLeastNumbers1(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	return core(arr, k, 0, len(arr)-1)
}

func core(arr []int, k int, left, right int) []int {
	if left < right {
		middleIndex := partion(arr, left, right)
		if middleIndex == k {
			return arr[:middleIndex]
		} else if middleIndex > k {
			return core(arr, k, left, middleIndex-1)
		} else {
			return core(arr, k, middleIndex+1, right)
		}
	}
	return arr[:k]
}

func partion(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
