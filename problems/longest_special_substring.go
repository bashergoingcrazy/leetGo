package problems

func maximumLength(s string) int {
    n := len(s) 
    check := func (x int) bool {
        w := x + 2
        arr := make([]int, 26)
        for _, c := range s {
            arr[c-'a'] ++
            if w >= arr[c-'a'] {
                return true
            }
        }
        return false
    }
    l, r := 0, n-1
    for l + 1< r {
        mid := l + (r-l)/2
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }
    if l == 0 {
        return -1
    } 
    return l
}

// func isSpecial(s string) bool {
//     arr := make([]int, 26)
//     for _, r := range s {
//         arr[r-'a'] = 1 
//     }
//     ct := 0
//     for _, val := range arr {
//         ct += val
//         if ct > 1 {
//             return false
//         }
//     }
//     return true
// }
//
// func maximumLength(s string) int {
//     mp := make(map[string]int)
//     n := len(s)
//
//     for i:=0; i<n; i++ {
//         for j:=i+1; j<=n; j++ {
//             mp[s[i:j]]++
//         }
//     }
//     
//     maxLength := -1 
//     for key, val := range mp {
//         if val >= 3  && isSpecial(key){
//             maxLength = max(maxLength, len(key))
//         }
//     }
//     return maxLength
// }
