package _9_跳跃游戏

// 判断每一点的最大跳跃距离内是否有0这个元素
func canJump(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return true
	}
	for i := len(nums)-2; i >= 0;i-- {
		if  nums[i] == 0 {
			j := i - 1
			for  ; j >= 0; j-- {
				if nums[j] > i - j {
					i = j
					break
				}
			}
			if j == -1{
				return false
			}
		}
	}
	return true
}
