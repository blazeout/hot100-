package _1_最小栈

type MinStack struct {
	stack1 []int // 正常栈
	stack2 []int // 辅助栈
}


func Constructor() MinStack {
	return MinStack{}
}

// 当一个元素要入栈时，我们取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中；

// 当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出；

// 在任意一个时刻，栈内元素的最小值就存储在辅助栈的栈顶元素中。

func (s *MinStack) Push(val int)  {
	if len(s.stack1) == 0 {
		s.stack1 = append(s.stack1, val)
		s.stack2 = append(s.stack2, val)
	}else {
		min := s.stack2[len(s.stack2)-1]
		if val < min {
			s.stack2 = append(s.stack2, val)
		}else {
			s.stack2 = append(s.stack2, min)
		}
		s.stack1 = append(s.stack1, val)
	}
}


func (s *MinStack) Pop()  {
	s.stack1 = s.stack1[:len(s.stack1)-1]
	s.stack2 = s.stack2[:len(s.stack2)-1]
}


func (s *MinStack) Top() int {
	value := s.stack1[len(s.stack1)-1]
	return value
}


func (s *MinStack) GetMin() int {
	value := s.stack2[len(s.stack2)-1]
	return value
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
