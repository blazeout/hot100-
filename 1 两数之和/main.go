package __两数之和

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{} // [int]表示下标存入for range 出来的值
	for i,v := range nums{
		p,ok:= hashTable[target - v]
		if ok{
			return []int{p,i}
		}
		hashTable[v] = i
	}
	return nil
}
