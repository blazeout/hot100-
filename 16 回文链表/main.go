package _6_回文链表

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历链表,将链表元素加入数组,再在数组中判断是否为回文 O(n), O(n)
func isPalindrome1(head *ListNode) bool {
    if head == nil {
        return false
    }
    nums := make([]int, 0)
    for head != nil {
        nums = append(nums , head.Val)
        head = head.Next
    }
    ok := isOK(nums)
    return ok
}

func isOK(nums []int) bool {
    for i := 0; i < len(nums); i++ {
        if nums[i] != nums[len(nums)-1-i] {
            return false
        }
    }
    return true
}

//避免使用O(n)的额外空间的方法就是改变输入

func endOfFirstHalf(head *ListNode) *ListNode {
	// 快指针走两步慢指针走一步当快指针走到了链表末尾时,慢指针就走到了中间部分
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	return pre
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 1. 找到前半部分链表的尾节点
	FirstNode := endOfFirstHalf(head)
	// 2. 反转后半部分链表
	retNode := reverseList(FirstNode.Next)
	// 3. 判断是否回文 当p2走到结尾时,则比较结束 忽略中间的节点 (当链表节点为奇数个时,中间节点算前面那部分的)
	p1 := head
	p2 := retNode
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 4. 恢复链表
	FirstNode.Next = reverseList(retNode)
	// 5. 返回结果
	return true
}
