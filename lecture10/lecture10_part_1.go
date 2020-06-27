package lecture10

/*
Problem:
	Maximum Profit in a Grid (homework for lecture9)

	A robot is located at the top-left corner of a m x n grid (marked 'S' in the diagram below).
	The robot can only move either down or right at any point in time.
	The robot is trying to reach the bottom-right corner of the grid (marked 'E' in the diagram below).
	Each cell contains a coin the robot can collect.

	What is the path that lead to the maximum profit the robot can accumulate?

	+---+---+---+---+
	| S | 2 | 2 | 1 |
	+---+---+---+---+
	| 3 | 1 | 1 | 1 |
	+---+---+---+---+
	| 4 | 4 | 2 | E |
	+---+---+---+---+
*/

// Time complexity: O(mn)
// Space complexity: O(mn)
func maxProfit(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = grid[i][j]
			if i > 0 && j > 0 {
				dp[i][j] += max(dp[i-1][j], dp[i][j-1])
			} else if i > 0 {
				dp[i][j] += dp[i-1][j]
			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return getPath(dp, m-1, n-1, [][]int{})
}

// getPath reconstructs the most profitable path
func getPath(dp [][]int, i, j int, path [][]int) [][]int {
	if i == 0 && j == 0 {
		return append(path, []int{i, j})
	} else if i == 0 {
		path = getPath(dp, i, j-1, path)
	} else if j == 0 {
		path = getPath(dp, i-1, j, path)
	} else {
		if dp[i-1][j] > dp[i][j-1] {
			path = getPath(dp, i-1, j, path)
		} else {
			path = getPath(dp, i, j-1, path)
		}
	}
	return append(path, []int{i, j})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
