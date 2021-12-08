package _6_最长的回文子串

func longestPalindrome(s string) string {
	// 回文子串 : 正着和反着一样
	if s == "" {
		return ""
	}
	// 中心扩散法
	start, end := 0,0
	// 以1个字符为中心, 或者以2个字符为中心
	for i := 0; i < len(s); i++ {
		left1, right1 := isPalid(s, i, i)
		left2, right2 := isPalid(s, i, i+1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func isPalid(s string ,left,right int) (int, int) {
	for ;left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {}
	return left+1, right-1
}
