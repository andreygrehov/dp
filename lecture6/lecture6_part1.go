package lecture6

/*
Problem:
	Climbing Stairs (k steps, space optimized)

	You are climbing a stair case. It takes n steps to reach to the top.
	Each time you can climb 1..k steps.
	In how many distinct ways can you climb to the top?
*/

// Time complexity: O(nk)
// Space complexity: O(k)
func climbStairsKSteps(n int, k int) int {
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i % k] += dp[(i-j) % k]
		}
	}
	return dp[n % k]
}