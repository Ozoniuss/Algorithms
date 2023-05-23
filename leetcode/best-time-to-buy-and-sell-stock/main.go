package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

func maxProfitDP(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minIdx := 0
	mprofit := 0

	for i := 0; i < len(prices); i++ {
		// In this case we are certain it's not going to get a better profit.
		if prices[minIdx] > prices[i] {
			minIdx = i
			continue
		}
		if prices[i]-prices[minIdx] > mprofit {
			mprofit = prices[i] - prices[minIdx]
		}
	}
	return mprofit
}
func main() {

}
