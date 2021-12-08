package _0_删除链表的倒数第N个节点

// ListNode Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head
	// 让fast指针先走n步
	for n > 0 {
		fast = fast.Next
		n--
	}
	// 如果走到最后为空那么删除的节点就是头结点
	if fast == nil {
		return head.Next
	}
	// 要保证slow节点永远是删除节点的前一个,那么fast就要先走一步
	fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
