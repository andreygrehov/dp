package lecture15

/*
Video: https://youtu.be/g0VjciqYeDU

Problem:
	Coin change

	Given an unlimited supply of coins of given denominations,
	find the unique number of ways to make a change of size n.

	Denominations:
	coins = [1, 2, 3, 5]

	Transition function:
	i >= 1: f[i][1] = f[i-1][1]
	i >= 2: f[i][2] = f[i-1][1] + f[i-2][2]
	i >= 3: f[i][3] = f[i-1][1] + f[i-2][2] + f[i-3][3]
	i >= 5: f[i][5] = f[i-1][1] + f[i-2][2] + f[i-3][3] + f[i-5][5]
*/

/*

n = 3 coins = 1,2,3

No duplicates
for _, coin := range coins {
	for i := 1; i <= n; i++ {

coin = [1]

    (1)      (1)      (1)
3 ------ 2 ------ 1 ------ 0

coin = [1,2]

    (1)      (1)     (1)
/------ 2 ------ 1 ------ 0
3
\------- 1 ------ 0
    (2)     (1)

coin = [1,2,3]

    (1)     (1)      (1)
/------ 2 ------ 1 ------ 0
|
|  (3)
3 ------ 0
|
|
\------- 1 ------ 0
    (2)       (1)

Answer: 3

With duplicates
for i := 1; i <= n; i++ {
	for _, coin := range coins {


coins = [1]

     (1)
1 ------- 0


coins = [1,2]

    (1)       (1)
/------- 1 ------- 0
2
\--------0
    (2)

coins = [1,2,3]

             (1)      (1)
   (1)   /------- 1 ------- 0
/------ 2
3        \--------0
|              (2)
|  (2)     (1)
|----- 1 ----- 0
|
|  (3)
|----- 0

Answer: 4

*/

func coinChange(n int, coins []int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= n; i++ {
			if i-coin >= 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[n]
}

func coinChangeUniqueWays(n int, coins []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len(coins))
	}

	for i := range coins {
		dp[0][i] = 1
	}

	//i >= 1: f[i][1] = f[i-1][1]
	//i >= 2: f[i][2] = f[i-1][1] + f[i-2][2]
	//i >= 3: f[i][3] = f[i-1][1] + f[i-2][2] + f[i-3][3]
	//i >= 5: f[i][5] = f[i-1][1] + f[i-2][2] + f[i-3][3] + f[i-5][5]
	for i := 0; i <= n; i++ {
		for j := range coins {
			for k := 0; k <= j; k++ {
				if i-coins[k] < 0 {
					continue
				}
				dp[i][j] += dp[i-coins[k]][k]
			}
		}
		//if i >= 1 {
		//	dp[i][0] = dp[i-1][0]
		//	dp[i][1] = dp[i-1][0]
		//	dp[i][2] = dp[i-1][0]
		//	dp[i][3] = dp[i-1][0]
		//}
		//
		//if i >= 2 {
		//	dp[i][1] += dp[i-2][1]
		//	dp[i][2] += dp[i-2][1]
		//	dp[i][3] += dp[i-2][1]
		//}
		//
		//if i >= 3 {
		//	dp[i][2] += dp[i-3][2]
		//	dp[i][3] += dp[i-3][2]
		//}
		//
		//if i >= 5 {
		//	dp[i][3] += dp[i-5][3]
		//}
	}

	return dp[n][len(coins)-1]
}
