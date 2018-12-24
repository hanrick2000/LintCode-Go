/**
 * @param nums: an array of Integer
 * @param target: target = nums[index1] + nums[index2]
 * @return: [index1 + 1, index2 + 1] (index1 < index2)
 */
func twoSum (nums []int, target int) []int {
    left, right := 0, len(nums) - 1
    for left < right {
        if nums[left] + nums[right] == target {
            return []int{left + 1, right + 1}
        } else if nums[left] + nums[right] > target {
            right--
        } else {
            left++
        }
    }
    return []int{-1, -1}
}

