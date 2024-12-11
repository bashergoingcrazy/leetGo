package problems

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minSubArrayLen(target int, nums []int) int {
	fast, n := 0, len(nums)
	currentSum, res := 0, math.MaxInt32

	for i := 0; i < n; i++ {
		// Expand the window by moving `fast` until sum >= target
		for fast < n && currentSum < target {
			currentSum += nums[fast]
			fast++
		}

		// Check if the current window meets the condition
		if currentSum >= target {
			res = min(res, fast-i)
		}

		// Shrink the window from the left
		currentSum -= nums[i]
	}

	// If `res` was not updated, no valid subarray exists
	if res == math.MaxInt32 {
		return 0
	}
	return res
}