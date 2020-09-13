/*
283. 移动零
https://leetcode-cn.com/problems/move-zeroes/
*/

package main

func main() {

}

//双指针
func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	j:=0
	for i :=0;i<len(nums) ;i++  {
		if nums[i] !=0{
			nums[i] ,nums[j] =nums[j],nums[i]
			j++
		}
	}
}
