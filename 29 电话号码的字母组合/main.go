package _9_电话号码的字母组合

// 真的好难明天要再看一下



func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var mp = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combinations []string
	var dfs func(combination string, digits string)
	dfs = func(combination string, digits string) {
		if len(digits) == 0 {
			combinations = append(combinations, combination)
			return
		}
		for  _, v := range mp[string(digits[0])] {
			combination = combination+string(v)
			dfs(combination ,digits[1:])
			// 回溯
			combination = combination[:len(combination)-1]
		}
	}
	dfs("", digits)
	return combinations
}