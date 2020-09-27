/*
46. 全排列
https://leetcode-cn.com/problems/permutations/
*/
//2020/09/27

package main

func main() {

}

//回溯
var result [][]int

func back(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		temp := make([]int, len(nums))
		copy(temp, pathNums)
		result = append(result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			back(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	pathNums := []int{}
	used := make([]bool, len(nums))
	result = [][]int{}
	back(nums, pathNums, used)
	return result
}
