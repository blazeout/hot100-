package _1_括号生成

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(leftRemain int, rightRemain int, path string)
	dfs = func(leftRemain int, rightRemain int, path string){
		// 当path的长度与 2n相等那么就生成了一个解,将解加入res中
		if len(path) == 2 * n {
			res = append(res, path)
			return
		}
		// 如果左括号有剩余那么就要选择左括号
		if leftRemain > 0 {
			dfs(leftRemain-1, rightRemain, path+"(")
		}
		// 必须右括号的数量大于左括号的数量才可以选择右括号
		if leftRemain < rightRemain {
			dfs(leftRemain, rightRemain-1, path+")")
		}
	}
	dfs(n,n, "")
	return res
}
