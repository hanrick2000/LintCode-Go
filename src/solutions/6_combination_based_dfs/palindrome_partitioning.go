func partition(s string) [][]string {
    res := make([][]string, 0)
    candidate := make([]string, 0)
    if len(s) == 0 {
        return res
    }
    table := palindromeTable(s)
    dfs(s, 0, &candidate, &res, table)
    return res
}

func dfs(s string, cur int, candidate *[]string, res *[][]string, table [][]bool) {
    if cur == len(s) {
        partition := make([]string, len(*candidate))
        copy(partition, *candidate)
        *res = append(*res, partition)
        return
    }
    for i := cur; i < len(s); i++ {
        if table[cur][i] {
            *candidate =  append(*candidate, s[cur : i + 1])
            dfs(s, i + 1, candidate, res, table)
            *candidate = (*candidate)[0 : len(*candidate) - 1]
        }
    }
}

func palindromeTable(s string) [][]bool {
    n := len(s)
    res := make([][]bool, 0)
    for i := 0; i < n; i++ {
        res = append(res, make([]bool, n))
    }
    for i := 0; i < len(s); i++ {
        helper(s, i, i, res)
        helper(s, i, i + 1, res)
    }
    return res
}

func helper(s string, start int, end int, res [][]bool) {
    for start >= 0 && end < len(s) {
        if s[start] == s[end] {
            res[start][end] = true
            start--
            end++
        } else {
            break
        }
    }
} 
