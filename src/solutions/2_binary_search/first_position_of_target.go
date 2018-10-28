func binarySearch (nums []int, target int) int {
    // write your code here
    if nums == nil {
        return -1
    }
    start, end := 0, len(nums) - 1
    for ; start + 1 < end; {
        mid := start + (end - start) / 2
        if nums[mid] < target {
            start = mid
        } else if nums[mid] >= target {
            end = mid
        }
    }
    if nums[start] == target {
        return start
    } else if nums[end] == target {
        return end
    } else {
        return -1
    }
}

