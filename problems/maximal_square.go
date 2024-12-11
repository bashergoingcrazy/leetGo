package problems

func MaximalSquare(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    rowSums := make([][]int, m)
    colSums := make([][]int, m)

    // Initialize rowSums and colSums
    for i := 0; i < m; i++ {
        rowSums[i] = make([]int, n)
        colSums[i] = make([]int, n)
    }

    // Preprocess rowSums
    for i := 0; i < m; i++ {
        cnt := 0
        for j := n - 1; j >= 0; j-- {
            if matrix[i][j] == '1' {
                cnt++
            } else {
                cnt = 0
            }
            rowSums[i][j] = cnt
        }
    }

    // Preprocess colSums
    for j := 0; j < n; j++ {
        cnt := 0
        for i := m - 1; i >= 0; i-- {
            if matrix[i][j] == '1' {
                cnt++
            } else {
                cnt = 0
            }
            colSums[i][j] = cnt
        }
    }

    maxSize := 0

    // Check for the largest square
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' { // Only consider cells with 1
                size := 1 // Minimum square size
                for size <= min(rowSums[i][j], colSums[i][j]) {
                    // Check if bottom-right corner has sufficient row and col sums
                    if rowSums[i+size-1][j] >= size && colSums[i][j+size-1] >= size {
                        maxSize = max(maxSize, size)
                    } else {
                        break
                    }
                    size++
                }
            }
        }
    }

    return maxSize * maxSize
}

// Utility functions
// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

