func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    candidate := make([]int, 0)
    helper(nums, 0, &candidate, &res)
    return res
}

func helper(nums []int, cur int, candidate *[]int, res *[][]int) {
    temp := make([]int, len(*candidate))
    copy(temp, *candidate)
    *res = append(*res, temp)
    if cur == len(nums) {
        return
    }
    
    for i := cur; i < len(nums); i++ {
        if i != cur && nums[i] == nums[i - 1] {
            continue
        }
        *candidate = append(*candidate, nums[i])
        helper(nums, i + 1, candidate, res)
        *candidate = (*candidate)[0 : len(*candidate) - 1]
    }
}
