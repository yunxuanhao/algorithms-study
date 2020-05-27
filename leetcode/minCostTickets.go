package leetcode

// https://leetcode-cn.com/problems/minimum-cost-for-tickets/

func mincostTickets(days []int, costs []int) int {
	var lend = len(days)
	var maxDay = days[lend-1]
	var minDay = days[0]

	var dp = make(map[int]int)

	i := lend - 1
	for d := maxDay; d >= minDay; d-- {
		// i 表示 days 的索引
		// 也可提前将所有 days 放入 Set，再通过 set.contains() 判断
		if d == days[i] {
			dp[d] = min(dp[d+1]+costs[0], dp[d+7]+costs[1])
			dp[d] = min(dp[d], dp[d+30]+costs[2])
			i-- // 别忘了递减一天
		} else {
			dp[d] = dp[d+1] // 不需要出门
		}
	}

	return dp[minDay]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
