package _0_汉明距离

import "math/bits"

func hammingDistance(x int, y int) int {
    // 1. 依次取出这一位对应的二进制为即可
    arr1 := takeOutBit(x)
    arr2 := takeOutBit(y)
    count := 0
    for i := 0; i < 32; i++ {
        if arr1[i] != arr2[i] {
            count++
        }
    }
    return count
}

func takeOutBit(x int) []int{
    res := make([]int, 0)
    for i := 0; i < 32; i++ {
        ret := x & (1 << i)
        res = append(res, ret)
    }
    return res
}

// 不同的位数转换为异或之后计算1的个数


func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x^y))
}

func hammingDistance2(x int, y int) int {
	s := x^y
	ans := 0
	for s > 0 {
		ans += s & 1
		s = s >> 1
	}
	return ans
}

