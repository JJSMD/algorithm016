package main

func main() {
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums1)
	println(result)
}

//暴力 第一想法
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	flag := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == flag {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			flag = nums[i]
		}
	}

	return len(nums)
}

// 题解 双指针，快慢指针
func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
