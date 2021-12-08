package _4_无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	// 哈希集合, 记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针,初始值为-1, 代表此时还没有进入字符串,没有开始移动
	rightPoint, ans := -1, 0
	// for循环遍历string字符串
	for i := 0; i < n; i++ {
		if i != 0 {
			// i 代表左指针, 代表当 i 等于 0,1,2时
			delete(m, s[i-1])
		}
		// 此次for 循环控制右指针的循环前进直到碰到重复的子串
		for rightPoint + 1 < n && m[s[rightPoint+1]] == 0{
			m[s[rightPoint+1]]++
			rightPoint++
		}
		ans = max(ans, rightPoint - i + 1)
	}
	return ans
}

func max(a,b int) int {
	if a > b{
		return a
	}
	return b
}

func lengthOfLongestSubstring2(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}
	left, right, ans := 0,0,0
	m := make(map[byte]int)
	for right < strLen {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		}else {
			// 将左窗口更新为最近一次出现的位置+1
			if m[s[right]]+1 > left {
				left = m[s[right]]+1
			}
			m[s[right]] = right
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
