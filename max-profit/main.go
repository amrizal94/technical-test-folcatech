package main

import "fmt"

func MaxProfit(prices []int, i int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// buat array dp untuk menyimpan keuntungan maksimum yang dapat diperoleh dengan k transaksi
	// pada hari ke-i
	dp := make([][]int, i+1)
	for j := 0; j <= i; j++ {
		dp[j] = make([]int, n)
	}

	// lakukan iterasi untuk setiap transaksi
	for j := 1; j <= i; j++ {
		// simpan keuntungan maksimum yang dapat diperoleh pada hari ke-i
		maxDiff := -prices[0]
		for k := 1; k < n; k++ {
			// perhitungan keuntungan dengan melakukan transaksi pada hari ke-k
			dp[j][k] = max(dp[j][k-1], prices[k]+maxDiff)
			maxDiff = max(maxDiff, dp[j-1][k]-prices[k])
		}
	}

	return dp[i][n-1]
}

// fungsi helper untuk mengembalikan nilai maksimum dari dua bilangan
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{4, 11, 2, 20, 59, 80}
	i := 2
	fmt.Println(MaxProfit(prices, i)) // Output: 85
}
