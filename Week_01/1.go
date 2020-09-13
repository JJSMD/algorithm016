/*
1. 两数之和
https://leetcode-cn.com/problems/two-sum/
*/

package main

func main() {

}

//遍历
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if temp+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//哈希
func twoSum1(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}
