package lecture8

/*
Problem:
	Paid Staircase II

	You are climbing a paid staircase. It takes n steps to reach to the top and you have to
	pay p[i] to step on the i-th stair. Each time you can climb 1 or 2 steps.
	Return the cheapest path to the top of the staircase.

Template to reconstruct the path
================================

	path = []
	for curr = best_last_state; curr exists; curr = from[curr] {
		path.push(curr)
	}
	return path.reverse()

*/

// Time complexity: O(n)
// Space complexity: O(n)
func paidStaircase(n int, p []int) []int {
	dp := make([]int, n+1)
	from := make([]int, n+1)
	dp[0] = 0
	dp[1] = p[1]
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + p[i]
		if dp[i-1] < dp[i-2] {
			from[i] = i-1
		} else {
			from[i] = i-2
		}
	}

	path := []int{}
	for curr := n; curr >= 0; curr = from[curr] {
		path = append(path, curr)
		if curr == 0 {
			break
		}
	}

	return reverse(path)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}