/*
91. 解码方法
https://leetcode-cn.com/problems/decode-ways/
*/
//2020/10/25

package main

func main() {

}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	pre, cur, temp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			temp = cur
			cur = cur + pre
			pre = temp
		default:
			pre = cur
		}
	}
	return cur
}
