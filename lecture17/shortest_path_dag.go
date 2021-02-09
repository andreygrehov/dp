package lecture17

import (
	"math"

	"github.com/andreygrehov/dp/internal/util"
)

/*
Problem:
	Shortest Path in DAG

	Given a graph as an adjacency matrix, find the shortest path
	from the first to the last vertex.

	Objective function:
	F[i] is the shortest path from i to the last vertex.

	Transition function:
	F[i] = min[weight + F[j]]

	Base case:
	F[n] = 0
*/

func shortestPath(graph [][]int) int {
	if len(graph) == 0 {
		return 0
	}

	n := len(graph)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt64
	}

	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(graph[i]); j++ {
			weight := graph[i][j]
			if weight > 0 {
				dp[i] = util.Min(dp[i], weight+dp[j])
			}
		}
	}

	return dp[0]
}
