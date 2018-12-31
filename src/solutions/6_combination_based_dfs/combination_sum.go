func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := make([][]int, 0)
    candidate := make([]int, 0)
    helper(candidates, target, 0, &candidate, &res)
    return res
}

func helper(candidates []int, target int, cur int, 
            candidate *[]int, res *[][]int) {
    if target == 0 {
        temp := make([]int, len(*candidate))
        copy(temp, *candidate)
        *res = append(*res, temp)
        return
    }
    if target < 0 {
        return
    }
    for i := cur; i < len(candidates); i++ {
        *candidate = append(*candidate, candidates[i])
        helper(candidates, target - candidates[i], i, candidate, res)
        *candidate = (*candidate)[0 : len(*candidate) - 1]
    }
}
