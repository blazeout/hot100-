package _5_回文子串的个数


// 计算有多少个回文子串 :
// 1. 枚举出所有的子串,然后再判断是否回文
// O(n^3) O(1)
func countSubstrings1(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            if isPalindrome(s[i:j+1]) {
                count++
            }
        }
    }
    return count
}

func isPalindrome(s string) bool {
    i := 0
    j := len(s)-1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

// 2. 枚举每一个可能的回文中心, 然后用两个指针分别向左右两边扩展判断是否回文
//    中心可能以1个字符为中心, 或者以两个字符为中心

func countSubstrings2(s string) int {
	ans := 0
	// 此层for循环控制遍历整个字符串
	for i := 0; i < len(s); i++ {
		// 此层for循环控制回文中心
		for j := 0; j <= 1; j++ {
			left := i
			right := i + j
			for left >= 0 && right < len(s) && s[left] == s[right] {
				left--
				right++
				ans++
			}
		}
	}
	return ans
}

