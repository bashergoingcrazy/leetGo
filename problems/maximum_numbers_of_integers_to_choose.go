package problems 

func maxCount(banned []int, n int, maxSum int) int {
    m := make(map[int]bool)
    for k := range banned {
        m[k] = true
    }

    numChosen := 0
    currentSum := 0

    for i:=1; i<=n; i++ {
        if _,exists := m[i];  
    }
    
}
