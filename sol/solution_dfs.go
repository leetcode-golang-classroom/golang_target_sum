package sol

type NumRecord struct {
	end, total int
}

func findTargetSumWaysDFS(nums []int, target int) int {
	nLen := len(nums)
	cache := make(map[NumRecord]int)
	var dfsBackTrack func(end, total int) int
	// dfsBackTrack(end, total) : sum up from 0 to end with total's possible ways
	dfsBackTrack = func(end, total int) int {
		if end == nLen {
			if total == target {
				return 1 // found 1
			}
			return 0 // not found
		}
		if val, exists := cache[NumRecord{end: end, total: total}]; exists {
			return val
		}
		cache[NumRecord{end: end, total: total}] =
			dfsBackTrack(end+1, total+nums[end]) + dfsBackTrack(end+1, total-nums[end])
		return cache[NumRecord{end: end, total: total}]
	}
	return dfsBackTrack(0, 0)
}
