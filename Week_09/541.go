/*
541. 反转字符串 II
https://leetcode-cn.com/problems/reverse-string-ii/
*/
//2020/11/12

package main

func main() {

}

//暴力
func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i = i + 2*k {
		left := i
		right := i + k - 1
		if i+k-1 > len(s)-1 {
			right = len(s) - 1
		}
		for left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}
	return string(str)
}

//双指针
func reverseStr1(s string, k int) string {
	str := []byte(s)
	l := len(s)
	for i := 0; i < l; i += 2 * k {
		left := i
		right := l - 1
		if i+k-1 < l {
			right = i + k - 1
		}

		for left <= right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}
	return string(str)
}
