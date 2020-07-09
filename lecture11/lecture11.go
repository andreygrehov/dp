package lecture11

// this is NOT dynamic programming, but a regular recursion
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// this is a top-down dynamic programming (a.k.a. recursion + memoization)
func fibTopDown(n int) int {
	memo := map[int]int{}
	return fibTopDownHelper(n, memo)
}

func fibTopDownHelper(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = fibTopDownHelper(n-1, memo) + fibTopDownHelper(n-2, memo)
	return memo[n]
}


// this is a bottom-up dynamic programming (forward dynamic programming)
//
// f(i-1)
//      \
//       >-------> f(i)
//      /
// f(i-2)
//
func fibBottomUpDPForward(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// this is a bottom-up dynamic programming (backward dynamic programming)
//
//   -----> f(i+2)
//   |
// f(i)
//   |
//   -----> f(i+1)
//
func fibBottomUpDPBackward(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	for i := 1; i < n; i++ {
		dp[i+1] += dp[i] // dp[i] is already solved, use it to solve other subproblems
		dp[i+2] += dp[i]
	}
	return dp[n]
}