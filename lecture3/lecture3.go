package lecture3

/*

Problem:
	Find the sum of the first N numbers.

Objective function:
	f(i) is the sum of the first i elements.

Recurrence relation:
	f(n) = f(n-1) + n

 */

func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}