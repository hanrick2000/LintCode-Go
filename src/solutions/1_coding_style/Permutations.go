func permute(nums []int) [][]int {
    res := make([][]int, 0)
    if nums == nil || len(nums) == 0 {
        return res
    }
    n := len(nums)
    helper(&res, nums, 0, n)
    return res
}

func helper(res *[][]int, candidate []int, cur int, n int) {
    if (cur == n) {
        temp := make([]int, n)
        copy(temp, candidate)
        *res = append(*res, temp)
    }
    for i := cur; i < n; i++ {
        candidate[cur], candidate[i] = candidate[i], candidate[cur]
        helper(res, candidate, cur+1, n)
        candidate[i], candidate[cur] = candidate[cur], candidate[i]
    }
}
