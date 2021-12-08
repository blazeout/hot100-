package _2_下一个排列

func nextPermutation(nums []int)  {
	n := len(nums)
	i := n - 2
	// 1. 较小数 : 从后向前找第一个不满足nums[i] >= nums[i+1]的数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 2. 若是i == -1 ,那么即是已经为最大的降序数列
	if i >= 0 {
		// 3. 较大数 : 从后向前找第一个满足 j nums[i] < nums[j]
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	// 利用go语言语法糖交换数值
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}