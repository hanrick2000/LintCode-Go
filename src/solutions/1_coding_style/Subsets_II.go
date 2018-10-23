func subsetsWithDup(nums []int) [][]int {
    res := [][]int{}
    if nums == nil {
        return res
    }
    sort.Ints(nums)
    helper(&res, nums, &[]int{}, 0)
    return res
}

func helper(res *[][]int, nums []int, candidate *[]int, cur int) {
    n := len(nums)
    if cur >= n {
        temp := make([]int, len(*candidate))
        copy(temp, *candidate)
        *res = append(*res, temp)
        return
    }   
    if len(*candidate) == 0 || nums[cur] != (*candidate)[len(*candidate) - 1] {
        helper(res, nums, candidate, cur + 1)
    }
    *candidate = append(*candidate, nums[cur])
    helper(res, nums, candidate, cur + 1)
    *candidate = (*candidate)[0 : len(*candidate) - 1]
}
