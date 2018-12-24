/**
 * @param nums: an array of integer
 * @param target: an integer
 * @return: an integer
 */
import "sort"

func twoSum5 (nums []int, target int) int {
    // write your code here
    sort.Ints(nums)
    res := 0
    left, right := 0, len(nums) - 1
    for left < right {
        if nums[left] + nums[right] <= target {
            res += right - left
            left++
        } else {
            right--
        }
    }
    return res
}

// 2 7 11 15, 24
// left: 0, right: 3, count=3
// left: 1, right: 3, count=5
// left = 2, right = 3
// left = 3, right = 3
