package main

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums1, 3)

}

//暴力交换，类似插入操作 O(n)  k*n
func rotate(nums []int, k int) {
	for i := 1; i <= k; i++ {
		x := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = x
			x = temp
		}
	}
}

//新数组法 O(n)
func rotate2(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, a)
}

//三次翻转
func rotate3(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

}

func reverse(nums []int, begin int, end int) {
	temp := 0
	for begin < end {
		temp = nums[begin]
		nums[begin] = nums[end]
		nums[end] = temp
		begin++
		end--
	}
}

//切片组合
func rotate1(nums []int, k int) {
	k = k % len(nums)
	n := nums[len(nums)-k:]
	m := nums[0 : len(nums)-k]
	copy(nums, append(n, m...))
}
