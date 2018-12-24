/**
 * @param numbers: Give an array numbers of n integer
 * @return: Find all unique triplets in the array which gives the sum of zero.
 */
 import "sort"
 
func threeSum (numbers []int) [][]int {
    // write your code here
    nums := numbers
    sort.Ints(nums)
    n := len(nums)
    res := make([][]int, 0)
    
    for i := 0; i < n; i++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }
        j, k := i + 1, n - 1
        for j < k {
            if j != i + 1 && nums[j] == nums[j - 1] {
                j++
                continue
            }
            if nums[i] + nums[j] + nums[k] == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
            } else if nums[i] + nums[j] + nums[k] > 0 {
                k--
            } else {
                j++
            }
        }
    }
    return res
}

