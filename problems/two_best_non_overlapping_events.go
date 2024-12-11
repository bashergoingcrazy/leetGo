package problems


import (
	"sort"
)

func maxTwoEvents(events [][]int) int {
    sort.Slice(events, func(i, j int)bool {
        return events[i][0] < events[j][0]
    })

    n := len(events)
    f := make([]int,n+1)

    // create f such that for every index it has
    // max value of nodes coming after it
    for i:= n-1; i>=0; i-- {
        f[i] = max(f[i+1], events[i][2])
    }

    ans := 0
    for _, e := range events {
        v := e[2]
        left, right := 0, n
        for left < right {
            mid := left + (right-left)/2
            if events[mid][0] > e[1] {
                right = mid
            } else {
                left = mid + 1
            }
        }
        
        if left < n {
            v += f[left]
        }
        ans = max(ans, v)
    }
    return ans
}
