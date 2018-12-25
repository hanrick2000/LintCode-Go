func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    n := len(nums)
    sort.Ints(nums)
    
    for i := 0; i < n; i++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i + 1; j < n; j++ {
            if j != i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            left, right := j + 1, n - 1
            for left < right {
                if left != j + 1 && nums[left] == nums[left - 1] {
                    left++
                    continue
                }
                cur := nums[i] + nums[j] + nums[left] + nums[right]
                if cur == target {
                    res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    right--
                } else if cur < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return res
}
