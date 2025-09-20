package tasks

func BestProfit(prices []int) [2]int {
	var res = [2]int{-1, -1}

	if len(prices) < 2 {
		return res
	}

	bestProfit := 0
	minimalPrice := prices[0]
	minimalPriceIndex := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minimalPrice {
			minimalPrice = prices[i]
			minimalPriceIndex = i
		}
		
		profit := prices[i] - minimalPrice
		if profit > bestProfit {
			bestProfit = profit
			res[0] = minimalPriceIndex
			res[1] = i
		}
	}

	return res
}
