package _6_全排列

// 这种一般都是利用dfs深度遍历 + 回溯 + 剪枝 进行写
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	isVisited := map[int]bool{}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			newPath := make([]int, len(path))
			copy(newPath, path)
			res = append(res, newPath)
			return
		}
		for _, num := range nums {
			if isVisited[num] {
				continue
			}
			path = append(path, num)
			isVisited[num] = true
			dfs(path)
			path = path[:len(path)-1]
			isVisited[num] = false
		}
	}
	dfs([]int{})
	return res
}
