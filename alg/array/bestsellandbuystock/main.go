package main

import "fmt"

func maxProfit(prices []int) int {
	dayCnt := len(prices)
	if dayCnt == 0 || dayCnt == 1 {
		return 0
	}
	buy := 0
	sell := 1
	maxSell := 0
	totalProfit := 0
	profit := 0
	for sell < dayCnt {
		// If price of sell day if greater than that of buy day
		// we calculate the price and compare it with best profit so far
		if prices[buy] < prices[sell] && prices[sell] >= maxSell {
			diff := prices[sell] - prices[buy]
			if diff > profit {
				maxSell = prices[sell]
				profit = diff
			}
			sell++
			continue
		}
		// Otherwise, assign buy day to sell day
		// and sell day to next day of buy day
		totalProfit += profit
		profit = 0
		buy = sell
		sell = buy + 1
	}
	return totalProfit
}

func main() {
	days := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(days))
}
