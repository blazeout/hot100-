package main

import (
	"fmt"
)

func countBits(n int) []int {
	length := n + 1
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ret := countOne(i)
		ans[i] = ret
	}
	return ans
}

func countOne(n int) int {
	count := 0
	for n > 0 {
		n = n & (n-1)
		count++
	}
	return count
}

func main(){
	ret := countBits(2)
	fmt.Println(ret)
}
