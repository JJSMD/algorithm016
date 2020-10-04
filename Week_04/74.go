/*
74. 搜索二维矩阵
https://leetcode-cn.com/problems/search-a-2d-matrix/
*/
//2020/10/04

package main

func main() {

}

//组合成数组
func searchMatrix(matrix [][]int, target int) bool {

	nums := []int{}
	for i := 0; i < len(matrix); i++ {
		nums = append(nums, matrix[i]...)
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

//矩阵行列移动
func searchMatrix1(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) <= 0 {
		return false
	}
	rowSize, colSize := len(matrix), len(matrix[0])
	row, column := 0, colSize-1
	for row < rowSize && column >=0 {
		if matrix[row][column] == target {
			return true
		}
		if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}
