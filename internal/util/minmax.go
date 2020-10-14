package util

import "math"

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}
