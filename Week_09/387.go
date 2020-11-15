/*
387. 字符串中的第一个唯一字符
https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/
//2020/11/12
package main

func main() {

}

func firstUniqChar(s string) int {

	logs := make(map[uint8]int, 0)

	for i := 0; i < len(s); i++ {
		logs[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		val := logs[s[i]]
		if val == 1 {
			return i;
		}
	}
	return -1
}

