package main

// 将0移到后面,计算第一次出现0的位置
func moveZeroes(nums []int)  {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
