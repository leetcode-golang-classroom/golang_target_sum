package sol

func findTargetSumWaysV1(nums []int, target int) int {
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
	dp := make([]int, 2*sum+1)
	dp[sum+nums[0]] = 1
	dp[sum-nums[0]] += 1
	for end := 1; end < nLen; end++ {
		nextDp := make([]int, 2*sum+1)
		for total := -sum; total <= sum; total++ {
			if dp[total+sum] > 0 {
				nextDp[total+nums[end]+sum] += dp[total+sum]
				nextDp[total-nums[end]+sum] += dp[total+sum]
			}
		}
		copy(dp, nextDp)
	}
	return dp[sum+target]
}
