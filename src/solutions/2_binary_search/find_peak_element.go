func findPeakElement(nums []int) int {
    left, right := 0, len(nums) - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[mid - 1] {
            left = mid
        } else {
            right = mid
        }
    }
    if nums[left] > nums[right] {
        return left
    } else {
        return right
    }
}
