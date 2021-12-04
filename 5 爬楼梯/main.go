package __爬楼梯

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 2
	count := 0
	for i := 3; i <= n; i++ {
		count = a+b
		a = b
		b = count
	}
	return count
}
