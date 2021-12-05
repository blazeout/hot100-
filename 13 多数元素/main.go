package _3_多数元素


// 排序 O(nlogn)
// func majorityElement(nums []int) int {
//     sort.Ints(nums)
//     return nums[len(nums)/2]
// }

// 哈希表O(n)

// func majorityElement(nums []int) int {
//     m1 := make(map[int]int)
//     for  _,num:= range nums {
//         m1[num]++
//     }
//     for i := 0; i < len(nums);i++ {
//         if m1[nums[i]] > (len(nums)/2) {
//             return nums[i]
//         }
//     }
//     return 0
// }

// 摩尔投票法,众数x出现的次数一定比其他数字出现的次数之和还要多
// 设置一个vote变量,如果下一个数是众数,那么vote++,如果不是众数,那么vote--
// 每次进入循环时,就判断vote是否等于0,如果等于0,那么就将当前num设置为众数x
// 每当vote == 0时,就可以将前面的部分区间给舍弃掉这是因为剩余部分的总vote一定大于0
// 说明剩余部分的数组的众数一定不变
func majorityElement(nums []int) int{
	vote,x := 0,0
	for _,num := range nums{
		if vote == 0 {
			x = num
		}
		if num == x {
			vote++
		}else {
			vote--
		}
	}
	// 特别说明: 因为本题中说明了有个数字出现的次数超过长度的一半
	// 所以不用考虑是否存在众数的情况,若是题目没有说明一定存在众数
	// 那么就需要加入一个判定环节
	count := 0
	for _,num := range nums{
		if x == num {
			count++
		}
	}
	if count > len(nums)/2 {
		return x
	}
	return 0
}