/*
14. 最长公共前缀
https://leetcode-cn.com/problems/longest-common-prefix/
*/

package main

import "strings"

func main() {
	strs := []string{"dog", "racecar", "car"}
	longestCommonPrefix(strs)
}

//取第一个的第一个字符遍历，如果全部满足，在继续取第二个字符遍历，知道有出现不满足的情况
//以第一个元素为基础，向后匹配，不匹配则修改基础值，直到全部匹配，否则不存在

func longestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}
	if len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	x := strs[0]
	strs = strs[1:len(strs)]

	for _, curr := range strs {
		for strings.Index(curr, x) != 0 {
			if len(x) == 0 {
				return ""
			}
			x = x[0 : len(x)-1]
		}
	}
	return x
}

//二刷
func longestCommonPrefix1(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	flag := strs[0]
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i],flag) != 0{
			if len(flag) <= 1{
				return ""
			}
			flag=flag[0:len(flag)-1]
		}
	}
	return flag
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	flag := strs[0]
	for i := 0; i < len(flag); i++ {
		for j:=0;j<len(strs) ;j++  {
			if i>len(strs[j])-1{
				return flag[0:i]
			}
			if flag[i]!=strs[j][i] {
				return flag[0:i]
			}
		}
	}
	return flag
}