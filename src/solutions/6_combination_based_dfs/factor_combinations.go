func getFactors(n int) [][]int {
    res := make([][]int, 0)
    candidate := make([]int, 0)
    helper(2, n, &candidate, &res)
    return res
}

func helper(start int, n int, candidate *[]int, res *[][]int) {
    if start > n {
        return
    }

    for i := start; i <= n; i++ {
        if n % i == 0 {
            if n / i < i {
                break
            } else {
                *candidate = append(*candidate, i)
                toAdd := make([]int, len(*candidate))
                copy(toAdd, *candidate)
                toAdd = append(toAdd, n / i)
                *res = append(*res, toAdd)
                helper(i, n / i, candidate, res)
                *candidate = (*candidate)[0 : len(*candidate) - 1]
            }
        }
    }
}
