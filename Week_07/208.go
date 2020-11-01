/*
208. 实现 Trie (前缀树)
https://leetcode-cn.com/problems/implement-trie-prefix-tree/
*/
//2020/11/01

package main

func main() {

}

type Trie struct {
	name    string
	subTrie map[string]*Trie
	isWord  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		name:    "",
		subTrie: make(map[string]*Trie, 0),
		isWord:  false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, v := range word {
		value, ok := root.subTrie[string(v)]
		if ok {
			root = value
		} else {
			newNode := make(map[string]*Trie, 0)
			root.subTrie[string(v)] = &Trie{
				name:    string(v),
				subTrie: newNode,
				isWord:  false,
			}
			root = root.subTrie[string(v)]
		}
	}
	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if isWord, ok := this.searchbyWord(word); ok && isWord {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.searchbyWord(prefix)
	return ok
}

func (this *Trie) searchbyWord(word string) (bool, bool) {

	root := this
	var (
		value *Trie
		ok    bool
	)

	for _, v := range word {
		value, ok = root.subTrie[string(v)]
		if ok {
			root = value
			continue
		}
		return false, false
	}
	return value.isWord, true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
