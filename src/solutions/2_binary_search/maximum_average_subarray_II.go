func hasAverageLargerThan(nums []int, k int, target float64) bool {
    var curSum float64 = 0
    for i := 0; i < k; i++ {
        curSum += float64(nums[i]) - target
    }
    var prevSum float64 = 0
    curMin := prevSum
    for i := 0; i + k < len(nums); i++ {
        if curSum >= curMin {
            return true
        }
        curSum += float64(nums[i + k]) - target
        prevSum += float64(nums[i]) - target
        
        if curMin > prevSum {
            curMin = prevSum
        }
    }
    return curSum >= curMin
}
 
func findMaxAverage (nums []int, k int) float64 {
    var min, max float64 = float64(nums[0]), float64(nums[0])
    for i := 0; i < len(nums); i++ {
        if float64(nums[i]) < min {
            min = float64(nums[i])
        }
        if float64(nums[i]) > max {
            max = float64(nums[i])
        }
    }
    
    left, right := min, max
    for right - left > 0.000001 {
        mid := left + (right - left) / 2
        if hasAverageLargerThan(nums, k, mid) {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}

