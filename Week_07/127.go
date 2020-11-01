/*
127. 单词接龙
https://leetcode-cn.com/problems/word-ladder/description/
*/
//2020/10/07

package main

func main() {
	strs := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	ladderLength("hit", "cog", strs)
}

//BFS
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

//two-end BFS
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, word := range wordList {
		dict[word] = true
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	q1[beginWord] = true
	q2[endWord] = true

	l := len(beginWord)
	steps := 0

	for len(q1) > 0 && len(q2) > 0 {
		steps++
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}

		q := make(map[string]bool) //临时set
		for k := range q1 {
			chs := []rune(k)
			for i := 0; i < l; i++ {
				ch := chs[i]
				for c := 'a'; c <= 'z'; c++ {
					chs[i] = c
					t := string(chs)
					if _, ok := q2[t]; ok { //如果已经在尾端集合里，已经找到，返回
						return steps + 1
					}
					if _, ok := dict[t]; !ok { //不在词库里
						continue
					}
					delete(dict, t) //删除
					q[t] = true     //加入到临时队列
				}
				chs[i] = ch //单词复位，继续BFS
			}
		}
		q1 = q
	}
	return 0
}
