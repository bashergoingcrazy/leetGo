// https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/
package problems

import (
	"slices"
	"sort"
)

// O(n) solution
func _maximumBeauty(nums []int, k int) int {
    m := slices.Max(nums)
    arr := make([]int, m+2)
    for _, n := range nums {
        arr[max(0, n-k)]++
        arr[min(n+k+1, m+1)]--
    }

    count := 0
    ans := 0
    for _, x := range arr {
        count += x
        ans = max(count, ans)
    }
    return ans
}



// O(nlogn) solution
func maximumBeauty(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)

    check := func (beauty int) bool {
        for i:=0; i+beauty<n; i++ {
            // this diff will always be positive
            diff := nums[i+beauty] - nums[i]
            if diff <= 2*k {
                return true
            }
        }
        return false
    }

    l, r := 1, n-1
    res := 1
    for l<=r {
        mid := l + (r-l)/2
        if check(mid) {
            res = mid
            l = mid + 1
        } else {
            r = mid -1
        }
    }
    return res
}
