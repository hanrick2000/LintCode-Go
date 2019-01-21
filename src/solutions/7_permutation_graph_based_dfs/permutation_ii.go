func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    helper(nums, 0, &res)
    return res
}

func helper(nums []int, cur int, res *[][]int) {
    n := len(nums)
    
    if cur == n {
        toAdd := make([]int, n)
        copy(toAdd, nums)
        *res = append(*res, toAdd)
    }
    
    visited := make(map[int]struct{})
    for i := cur; i < n; i++ {
        if _, ok := visited[nums[i]]; !ok {
            visited[nums[i]] = struct{}{}
            nums[i], nums[cur] = nums[cur], nums[i]
            helper(nums, cur + 1, res)
            nums[i], nums[cur] = nums[cur], nums[i]
        }
    }
}
