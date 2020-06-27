package lecture10

/*
Problem:
	Paint Fence With Two Colors

	There is a fence with n posts, each post can be painted with either green or blue color.
	You have to paint all the posts such that no more than two adjacent fence posts have the same color.
	Return the total number of ways you can paint the fence.
*/

func numWays(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// green = 1
	// blue = 0
	dp[1][0] = 1
	dp[1][1] = 1
	dp[2][0] = 2 // 10, 00
	dp[2][1] = 2 // 01, 11

	for i := 3; i <= n; i++ {
		for j := 0; j <= 1; j++ {
			dp[i][j] = dp[i-1][1-j] + dp[i-2][1-j]
		}
	}

	return dp[n][0] + dp[n][1]
}
