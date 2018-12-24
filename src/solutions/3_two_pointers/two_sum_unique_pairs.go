/**
 * @param nums: an array of integer
 * @param target: An integer
 * @return: An integer
 */
import "sort"

func twoSum6 (nums []int, target int) int {
    // write your code here
    sort.Ints(nums)
    left, right := 0, len(nums) - 1
    res := 0
    
    for left < right {
        if left != 0 && nums[left] == nums[left - 1] {
            left++
            continue
        }
        if nums[left] + nums[right] == target {
            left++
            right--
            res++
        } else if nums[left] + nums[right] > target {
            right--
        } else {
            left++
        }
    }
    return res
}

