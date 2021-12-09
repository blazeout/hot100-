package _3_搜索旋转排序数组

func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			return mid
		}
		// 如果nums[mid] >= nums[left], 那么mid位于数组的左侧
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			}else {
				left = mid + 1
			}
		}else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			}else {
				right = mid - 1
			}
		}

	}
	return -1
}
