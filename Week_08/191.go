/*
191. 位1的个数
https://leetcode-cn.com/problems/number-of-1-bits/
*/
//2020/11/08

package main

func main() {

}

//位移
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

func hammingWeight1(num uint32) int {
	count := 0
	for num > 0 {
		num &= (num - 1)
		count++
	}
	return count
}
