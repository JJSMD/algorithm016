/*
21. 合并两个有序链表
https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/

package main

func main() {

}

//迭代 O(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	rel := new(ListNode)
	prev := rel

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return rel.Next
}

//递归 O(n+m)
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists1(l2, l1.Next)
		return l1
	}
}
