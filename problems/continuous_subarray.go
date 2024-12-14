package problems

// func abs(a int)int {
//     if a < 0 {
//         return -a
//     }
//     return a
// }

// O(n^2) solution
func _continuousSubarrays(nums []int) int64 {
    count := 0  
    for i := range nums {
        count++
        j := i 
        currMin := min(nums[j],nums[i])
        currMax := max(nums[j],nums[i])
        for currMax - currMin <= 2 {
            count++
            j++
            if j >= len(nums) {
                break
            } 
            currMin = min(nums[j],currMin)
            currMax = max(nums[j],currMax)
        }
    }
    
    return int64(count)
}

// O(n) solution using two pointers which i initally thought would be a waste of time
// but turns out it is just amortized solution i didn't thought the complexity of a medium
// would be this high that's why i abandoned my thought process in this approach when i 
// first saw the left pointer to backtrack in this sol 
func totalSubarray(n int) int {
    return n*(n+1)/2
}

func continuousSubarrays(nums []int) int64 {
    count := 0
    l ,r := 0,0
    minIndex ,maxIndex := 0,0
    for r < len(nums) {
        if nums[maxIndex] - nums[r] > 2 {
            count += totalSubarray(r-l)
            count -= totalSubarray(r - maxIndex - 1)
            l = maxIndex+1
            maxIndex , minIndex = l,l
            for k:= l+1; k<=r; k++ {
                if nums[k] >= nums[maxIndex] {
                    maxIndex = k
                }
                if nums[k] <= nums[minIndex] {
                    minIndex = k
                }
                
            }
        }

        if nums[r] - nums[minIndex] > 2 {
            count += totalSubarray(r-l)
            count -= totalSubarray(r - minIndex - 1)
            l = minIndex + 1
            maxIndex , minIndex = l,l
            for k:= l+1; k<=r; k++ {
                if nums[k] >= nums[maxIndex] {
                    maxIndex = k
                }
                if nums[k] <= nums[minIndex] {
                    minIndex = k
                }
                
            }
        }

        if nums[r] >= nums[maxIndex] {
            maxIndex = r
        }
        if nums[r] <= nums[minIndex] {
            minIndex = r
        }
        r++
    }
    count += totalSubarray(r-l)
    return int64(count)
}
