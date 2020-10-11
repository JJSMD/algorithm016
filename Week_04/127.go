/*
127. 单词接龙
https://leetcode-cn.com/problems/word-ladder/description/
*/
//2020/10/07

package main

func main() {
	strs:=[]string{"hot","dot","dog","lot","log","cog"}

	ladderLength("hit","cog",strs)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordMap := make(map[string]bool)
	isHas := false
	for _, v := range wordList {
		wordMap[v] = true
		if v == endWord {
			isHas = true
		}
	}
	if !isHas {
		return 0
	}
	queue := []string{}
	queue = append(queue, beginWord)
	l := len(beginWord)
	steps := 0

	for len(queue) > 0 {
		steps++
		size := len(queue)
		for i := size; i > 0; i-- {
			s := queue[0]
			queue = queue[1:]
			chs := []rune(s)
			for j := 0; j < l; j++ {
				ch := chs[j]
				for c := 'a'; c <= 'z'; c++ {
					if c == ch {
						continue
					}
					chs[j] = c
					t := string(chs)
					if t == endWord {
						return steps + 1
					}
					if _, ok := wordMap[t]; !ok {
						continue
					}
					delete(wordMap, t)
					queue = append(queue, t)
				}
				chs[j] = ch
			}
		}
	}

	return 0
}
