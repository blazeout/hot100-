package _4_在排序数组中查找元素的第一个和最后一个位置


// 二分查找
func searchRange1(nums []int, target int) []int {
	// 找左边界, 和右边界
	// 1. 找第一个 >= target 的数
	left := searchInt(nums , target)
	// 2. 找第一个 > target 的数
	right := searchInt(nums, target+1)
	if left == len(nums) || nums[left] != target {
		return []int{-1,-1}
	}
	return []int{left, right-1}
}

func searchInt(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		}else {
			left = mid + 1
		}
	}
	return left
}
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target && nums[r] == target {
			return []int{l, r}
		}
		if nums[l] == target {
			r--
		} else if nums[r] == target {
			l++
		} else {
			l++
			r--
		}
	}
	return []int{-1, -1}
}

