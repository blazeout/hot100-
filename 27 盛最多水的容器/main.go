package _7_盛最多水的容器



// 面积公式 ： S = min(a[i], a[j]) * (j-i)
// 暴力法 O(n^2) 超出时间限制
// 双指针法左指针指向开始, 右指针指向结尾
// 面积取决于短板。①因此即使长板往内移动时遇到更长的板，矩形的面积也不会改变；遇到更短的板时，面积会变小。
// ②因此想要面积变大，只能让短板往内移动(因为移动方向固定了)，当然也有可能让面积变得更小，但只有这样才存在让面积变大的可能性
func maxArea(height []int) int {
	length := len(height)
	left := 0
	right := length-1
	ans := 0
	for left < right {
		ans = max(ans, min(height[left], height[right]) * (right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
