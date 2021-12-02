package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	stack := make([]byte,0)
	for i := 0; i < n; i++ {
		// parids[s[i]] > 0 说明就是碰到了右括号
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main(){
	s := "[()]"
	ret := isValid(s)
	fmt.Println(ret)
}
