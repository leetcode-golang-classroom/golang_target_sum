package sol

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	if abs(target) > sum {
		return 0
	}
	nLen := len(nums)
	dp := make([][]int, nLen)
	for row := range dp {
		dp[row] = make([]int, 2*sum+1)
	}
	dp[0][sum+nums[0]] = 1
	dp[0][sum-nums[0]] += 1
	for end := 1; end < nLen; end++ {
		for total := -sum; total <= sum; total++ {
			if dp[end-1][total+sum] > 0 {
				dp[end][total+nums[end]+sum] += dp[end-1][total+sum]
				dp[end][total-nums[end]+sum] += dp[end-1][total+sum]
			}
		}
	}
	return dp[nLen-1][sum+target]
}
