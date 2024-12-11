package problems



func abs(x int) int {
    if x < 0 { return -x}
    return x
}


func isArraySpecial(nums []int, queries [][]int) []bool {
    n := len(nums)
    ct := make([]int, n) 
    x := 0
    for i:=1; i<n; i++ {
        if abs(nums[i]-nums[i-1])%2==0 {
            x++
        }
        ct[i]=x
    }

    res := make([]bool, len(queries)) 
    for i,q := range queries {
        if ct[q[0]] == ct[q[1]] {res[i]=true}
    }
    return res
}
