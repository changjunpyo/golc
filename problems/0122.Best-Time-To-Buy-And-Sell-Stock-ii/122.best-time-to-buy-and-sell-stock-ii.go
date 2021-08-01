package problems

func maxProfit(prices []int) int {
	min := prices[0]
	total := 0
	for i:=1; i<len(prices); i++{
		if prices[i-1] > prices[i] {
			total += prices[i-1] - min
			min = prices[i]
		} else if min > prices[i]{
			min = prices[i]
		} else if i == len(prices)-1 && min < prices[i]{
			total += prices[i]- min
		}
	}
	return total
}
