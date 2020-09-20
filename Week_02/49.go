/*
242. 有效的字母异位词
https://leetcode-cn.com/problems/group-anagrams/
*/

package main

import (
	"encoding/json"
	"sort"
	"strings"
)

func main() {

}

//排序
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	group := make(map[string]int)
	result := make([][]string, 0)
	index := 0
	for _, v := range strs {
		vArr := strings.Split(v, "")
		sort.Strings(vArr)
		toStr := strings.Join(vArr, "")
		if idx, ok := group[toStr]; !ok {
			group[toStr] = index
			result = append(result, []string{v})
			index++
		} else {
			result[idx] = append(result[idx], v)
		}
	}
	return result
}

//计数
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		temp := make([]int, 26)

		for _, c := range str {
			temp[c-'a'] ++
		}
		bytes, _ := json.Marshal(temp)
		m[string(bytes)] = append(m[string(bytes)], str)
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
