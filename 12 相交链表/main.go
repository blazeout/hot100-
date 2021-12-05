package _2_相交链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	point1 := headA
	point2 := headB
	for point1 != point2 {
		if point1 != nil {
			point1 = point1.Next
		}else {
			point1 = headB
		}
		if point2 != nil {
			point2 = point2.Next
		}else {
			point2 = headA
		}
	}
	return point1
}
