func combine(n int, k int) [][]int {
    res := make([][]int, 0)
    candidate := make([]int, 0)
    helper(n, k, 1, &candidate, &res)
    return res
}

func helper(n int, k int, cur int, candidate *[]int, res *[][]int) {
    if len(*candidate) == k {
        temp := make([]int, len(*candidate))
        copy(temp, *candidate)
        *res = append(*res, temp)
        return
    } 
    if cur > n {
        return
    }
    helper(n, k, cur + 1, candidate, res)
    *candidate = append(*candidate, cur)
    helper(n, k, cur + 1, candidate, res)
    *candidate = (*candidate)[0 : len(*candidate) - 1]
}
