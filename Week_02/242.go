/*
242. 有效的字母异位词
https://leetcode-cn.com/problems/valid-anagram/
*/

package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	words := make(map[rune]int)

	for _, v := range s {
		words[v]++
	}
	for _, v := range t {
		words[v]--
		if words[v] == 0 {
			delete(words, v)
		}
	}

	if len(words) == 0 {
		return true
	}
	return false
}
