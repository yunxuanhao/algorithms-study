package leetcode

// https://leetcode-cn.com/problems/maximal-square/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 {
		return 0
	}
	max := 0
	iMax, jMax := len(matrix), len(matrix[0])
	dp := make([][]int, iMax)
	for i, _ := range dp {
		dp[i] = make([]int, jMax)
	}
	for i := 0; i < iMax; i++ {
		for j := 0; j < jMax; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minN(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max
}

func minN(nums ...int) int {
	min := 10000000
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}
