package leetcode

// Returns the max profit you can make
func MaxProfit(prices []int) int {
	profit := 0
	buy := prices[0]

	// Get profit
	for day := 1; day < len(prices); day++ {
		if prices[day] < buy {
			buy = prices[day]
		} else if profit < prices[day]-buy {
			profit = prices[day] - buy
		}
	}

	return profit
}

// In this function, you can re-buy after selling a stock
func MaxProfitII(prices []int) int {
	profit := 0
	buy := prices[0]

	// Get profit
	for day := 1; day < len(prices); day++ {
		// If sell increasing (or stays the same), next
		// If sell decrases -> sell at day-1 (if we sold at the first day, it would just be 0)
		if prices[day-1] > prices[day] {
			profit += max(prices[day-1]-buy, 0)
			buy = prices[day]
		}
	}

	// Either selling or not on the last day
	if prices[len(prices)-1] > buy {
		profit += max(prices[len(prices)-1]-buy, 0)
	}

	return profit
}

// NOTES
// Look up Kadane's Algorithm
// - We can assume that the lowest number is when we want to buy
// - for example [7, 1, 5, 3, 6]: 1 would always be the buy day
