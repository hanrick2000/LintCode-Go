/**
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func findMin (nums []int) int {
    left, right := 0, len(nums) - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[len(nums) - 1] {
            left = mid
        } else {
            right = mid
        }
    }
    if nums[left] < nums[right] {
        return nums[left]
    } else {
        return nums[right]
    }
}
