package _9_找到数组中消失的数字


// 哈希表
func findDisappearedNumbers(nums []int) []int {
    hashmap := map[int]int{}
    for _ , num := range nums {
        hashmap[num] = 1
    }
    res := make([]int,0)
    for i := 1; i <= len(nums); i++ {
        if hashmap[i] != 1 {
            res = append(res, i)
        }
    }
    return res
}
// 将所有正数作为数组下标，置对应数组值为相反值。那么，仍为正数的位置即为（未出现过）消失的数字。
func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

