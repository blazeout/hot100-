package __最大子序和
func maxSubArray(nums []int) int {
	// for range nums
	// 如果count < 0 说明这个元素对于子数组是负效益,需要舍弃整个前面子数组的值,并且更新为当前元素的值
	if len(nums) == 1 {
		return nums[0]
	}
	// count用来记录每个不同段的连续子数组
	count := 0
	// res 用来保存比较之后的最大值
	res := nums[0]
	for _, value := range nums {
		if count > 0 {
			count += value
		}else {
			count = value
		}
		res = max(res,count)
	}
	return res
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
