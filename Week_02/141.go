/*
141. 环形链表
https://leetcode-cn.com/problems/linked-list-cycle/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china
*/

package main

func main() {

}

//双指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//哈希
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	list := make(map[*ListNode]bool)
	curr := head
	for curr != nil {
		if _, ok := list[curr]; ok {
			return true
		}
		list[curr] = true
		curr = curr.Next
	}
	return false
}
