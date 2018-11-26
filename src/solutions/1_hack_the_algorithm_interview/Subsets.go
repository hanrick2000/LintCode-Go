func subsets(nums []int) [][]int {
    res := [][]int{}
    if nums == nil {
        return res
    }
    candidate := []int{}
    helper(&res, nums, &candidate, 0)
    return res
}

func helper(res *[][]int, nums []int, candidate *[]int, cur int) {
    if cur >= len(nums) {
        temp := make([]int, len(*candidate))
        copy(temp, *candidate)
        *res = append(*res, temp)
        return
    }
    helper(res, nums, candidate, cur + 1)
    *candidate = append(*candidate, nums[cur])
    helper(res, nums, candidate, cur + 1)
    *candidate = (*candidate)[0 : len(*candidate) - 1]
}

