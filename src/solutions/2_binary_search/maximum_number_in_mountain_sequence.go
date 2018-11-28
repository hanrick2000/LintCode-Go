/**
 * @param nums: a mountain sequence which increase firstly and then decrease
 * @return: then mountain top
 */
func mountainSequence (nums []int) int {
    // write your code here
    if len(nums) == 1 {
        return nums[0]
    }
    left, right := 0, len(nums) - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if nums[mid] < nums[mid + 1] {
            left = mid
        } else {
            right = mid
        }
    }
    if nums[left] > nums[right] {
        return nums[left]
    } else {
        return nums[right]
    }
}

