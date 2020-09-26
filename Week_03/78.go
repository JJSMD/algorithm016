/*
78. 子集
https://leetcode-cn.com/problems/subsets/
*/
//2020/9/26

package main

func main() {

}

//回溯
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)

	result = append(result, item)
	generatel(0, nums, &item, &result)
	return result
}

func generatel(index int, nums []int, item *[]int, result *[][]int) {
	if index >= len(nums) {
		return
	}

	*item = append(*item, nums[index])
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v
	}
	*result = append(*result, temp)
	//放入下一元素
	generatel(index+1, nums, item, result)
	//不放下一元素
	*item = (*item)[:len(*item)-1]
	generatel(index+1, nums, item, result)
}

//迭代，穷举二进制表示法
//用0表示元素不在子集，1表示在子集里，迭代出元素总长度n的所有0/1组合情况就可以得出所有子集,情况总数是（2n次幂-1）
func subsets1(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		res = append(res,append([]int(nil),set...))
	}
	return  res
}
