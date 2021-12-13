package _5_组合总和

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTemp := make([]int,len(temp))
				copy(newTemp, temp)
				res = append(res, newTemp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum + candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
