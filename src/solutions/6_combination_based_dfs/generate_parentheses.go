func generateParenthesis(n int) []string {
    candidate := make([]byte, 0)
    res := make([]string, 0)
    helper(0, n, &candidate, &res)
    return res
}

func helper(numClosed int, n int, candidate *[]byte, res *[]string) {
    if len(*candidate) == 2 * n {
        if numClosed == 0 {
            *res = append(*res, string(*candidate))
        }
        return
    }
    
    // Add ')' if allowed
    if numClosed > 0 {
        *candidate = append(*candidate, ')')
        helper(numClosed - 1, n, candidate, res)
        *candidate = (*candidate)[:len(*candidate) - 1]
    }
    
    // Add '('
    *candidate = append(*candidate, '(')
    helper(numClosed + 1, n, candidate ,res)
    *candidate = (*candidate)[:len(*candidate) - 1]
}
