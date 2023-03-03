func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indegrees := make([]int, n)
	for i := 0; i < n; i++ {
		j := favorite[i]
		indegrees[j]++
	}

	var queue []int
	for i := 0; i < n; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	dp := make([]int, n) // dp[i] is the longest path leading to i exclusively.
	var i int
	for len(queue) != 0 {
		i, queue = queue[0], queue[1:] // remove front
		j := favorite[i]
		dp[j] = max(dp[j], dp[i]+1)

		indegrees[j]--
		if indegrees[j] == 0 {
			queue = append(queue, j)
		}
	}

	res, res2 := 0, 0
	for i = 0; i < n; i++ {
		if indegrees[i] != 0 {
			length := 0
			for j := i; indegrees[j] > 0; j = favorite[j] {
				indegrees[j] = 0
				length++
			}

			if length == 2 {
				res2 += 2 + dp[i] + dp[favorite[i]]
			} else {
				res = max(res, length)
			}
		}
	}

	return max(res, res2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}