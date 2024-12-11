package problems

func minimumSize(nums []int, maxOperations int) int {
    isPossible := func (maxBagSize int) bool {
        count := 0
        for _, val := range nums {
            count += (val-1)/maxBagSize 
            if count > maxOperations {
                return false
            }
        }
        return true
    }
    left, right := 1, 0
    for _, val := range nums {
        if val > right {
            right = val
        }
    }

    for left < right {
        mid := left + (right-left)/2
        if isPossible(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
