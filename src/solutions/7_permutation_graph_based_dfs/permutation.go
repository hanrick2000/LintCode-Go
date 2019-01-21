func permute(nums []int) [][]int {
    res := make([][]int, 0)
    helper(nums, 0, &res)
    return res
}

func helper(nums []int, cur int, res *[][]int) {
    if cur == len(nums) {
        toAdd := make([]int, len(nums))
        copy(toAdd, nums)
        *res = append(*res, toAdd)
    }
    
    for i := cur; i < len(nums); i++ {
        nums[i], nums[cur] = nums[cur], nums[i]
        helper(nums, cur + 1, res)
        nums[i], nums[cur] = nums[cur], nums[i]
    }
}
