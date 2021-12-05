package __买卖股票的最佳时机

// 循环遍历数组
// 1.记录当前时刻的最小值
// 2.计算出这一天的价格减去最小值, 即这一天的最大利润
// 3.用这一天的最大利润和之前的最大利润相比得出总的最大利润
func maxProfit(prices []int) int {
	maxprofit := 0
	min := prices[0]
	nowProfit := 0
	for _, price := range prices {
		if price < min {
			min = price
		}else {
			nowProfit = price - min
		}
		maxprofit = Max(nowProfit, maxprofit)
	}
	return maxprofit
}

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
