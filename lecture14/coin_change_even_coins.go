package lecture14

/*
Problem:
	Coin change

	Given an unlimited supply of coins of given denominations,
	find the total number of ways to make a change of size n, by
	using an even number of coins.

	// 1, 3, 5, 10
	f[i][0] = f[i-1][1] + f[i-3][1] + f[i-5][1] + f[i-10][1]
	f[i][1] = f[i-1][0] + f[i-3][0] + f[i-5][0] + f[i-10][0]
*/

func coinChange(n int, coins []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = 1

	for i := 1; i <= n; i++ {
		//if i-1 >= 0 {
		//	dp[i][0] += dp[i-1][1]
		//	dp[i][1] += dp[i-1][0]
		//}
		//
		//if i-3 >= 0 {
		//	dp[i][0] += dp[i-3][1]
		//	dp[i][1] += dp[i-3][0]
		//}
		//
		//if i-5 >= 0 {
		//	dp[i][0] += dp[i-5][1]
		//	dp[i][1] += dp[i-5][0]
		//}
		//
		//if i-10 >= 0 {
		//	dp[i][0] += dp[i-10][1]
		//	dp[i][1] += dp[i-10][0]
		//}
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i][0] += dp[i-coin][1]
			dp[i][1] += dp[i-coin][0]
		}
	}

	return dp[n][1]
}
