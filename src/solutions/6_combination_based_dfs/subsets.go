func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    candidate := make([]int, 0)
    helper(nums, 0, &candidate, &res)
    return res
}

func helper(nums []int, cur int, candidate *[]int, res *[][]int) {
    if cur == len(nums) {
        temp := make([]int, len(*candidate))
        copy(temp, *candidate)
        *res = append(*res, temp)
        return
    }
    helper(nums, cur + 1, candidate, res)
    *candidate = append(*candidate, nums[cur])
    helper(nums, cur + 1, candidate, res)
    *candidate = (*candidate)[0 : len(*candidate) - 1]
}
