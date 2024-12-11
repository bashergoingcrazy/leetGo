package problems

func combinationSum3(k, n int) [][]int {
    var res [][]int
    var tmp []int

    var solve func (index, currentSum int)
    solve = func (index, currentSum int) {
        // base case 
        if len(tmp) == k {
            if currentSum == n {
                combination := make([]int, len(tmp))
                copy(combination,tmp)
                res = append(res, combination)
            }
            return
        }

        // prune conditions
        if currentSum > n || index > 9 {
            return
        }

        // recursive condition
        for i:=index; i<9; i++ {
            tmp = append(tmp, i)
            solve(i+1, currentSum+i)
            tmp = tmp[:len(tmp)-1] // backtrack
        }
    }
    solve (1,0)
    return res
}
