package problems

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    bucket := make(map[int]int)

    getBucketById := func (num int) (id int)  {
        if num >= 0 {
            id = num/(valueDiff+1)
        } else {
            id = (num/(valueDiff+1))-1
        }
        return
    }

    for l:=0; l<len(nums); l++ {
        id := getBucketById(nums[l])

        if _, exists := bucket[id];exists {
            return true
        }

        if val,exists := bucket[id-1]; exists && nums[l]-val <= valueDiff {
            return true
        }

        if val,exists := bucket[id+1]; exists && val - nums[l] <= valueDiff {
            return true
        }

        bucket[id] = nums[l]
        if l>=indexDiff {
            delete(bucket, getBucketById(nums[l-indexDiff]))
        }
    }
    return false
}
