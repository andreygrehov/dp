package lecture12

/*
Problem:
	Coin change

	Given an unlimited supply of coins of given denominations,
	find the total number of ways to make a change of size n.

	Transition function: f(n) = f(n-d_1) + f(n-d_2) + f(n-d_3) + ... + f(n-d_k),
	where d_1, d_2, d_3, ..., d_k are provided coin denomations.
*/

// coinChange(4)
func coinChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= 1 {
			dp[i] += dp[i-1]
		}
		if i >= 3 {
			dp[i] += dp[i-3]
		}
		if i >= 5 {
			dp[i] += dp[i-5]
		}
		if i >= 10 {
			dp[i] += dp[i-10]
		}
	}
	return dp[n]
}

func coinChangeWithDenomations(n int, coins []int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[n]
}