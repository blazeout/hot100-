package _0环形链表

type ListNode struct {
	Val int
	Next *ListNode
}

// 定义两个快慢指针,fast每次走两步,slow每次走一步,若是有环那么一定会相遇所以循环条件设定为fast != slow
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
