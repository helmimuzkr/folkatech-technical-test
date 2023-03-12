package main

func MaxProfit(prices []int, i int) int {
	var nProfit, nTransactions int
	n := len(prices)

	// Do iteration everyday to find the profit.
	for day := 0; day < n; day++ {
		var max int
		for j := day + 1; j < n; j++ {
			// Check if tommorow stock share is greater than today purchased share,
			// or smaller than the previous day share.
			if prices[j] < prices[day] || prices[j] < prices[j-1] {
				day = j - 1 // continue tomorrow
				break
			}

			// Finding the max profit.
			profit := prices[j] - prices[day]
			if profit > max {
				max = profit
			}

			// Checking the days limit
			if j == n-1 {
				day = j
			}
		}

		// Sum up the highest profit with the previous profit and check how many transactions have been made.
		if max != 0 {
			nProfit += max
			nTransactions++
			if nTransactions == i {
				break
			}
		}
	}

	return nProfit
}
