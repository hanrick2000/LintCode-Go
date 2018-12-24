/**
 * @param nums: The integer array you should partition
 * @param k: An integer
 * @return: The index after partition
 */
func partitionArray (nums []int, k int) int {
    prev := 0
    for i, val := range nums {
        if val < k {
            nums[i], nums[prev] = nums[prev], nums[i]
            prev++
        }
    }
    return prev
}

