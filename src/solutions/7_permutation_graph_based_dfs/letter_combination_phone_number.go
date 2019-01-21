func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    letterMap := []string{"", "", "abc", "def", "ghi",
                          "jkl", "mno", "pqrs", "tuv", "wxyz"}
    res := make([]string, 0)
    candidate := make([]byte, 0)
    helper(digits, 0, &candidate, &res, letterMap)
    return res
}

func helper(digits string, cur int, candidate *[]byte,
            res *[]string, letterMap []string) {
    if cur == len(digits) {
        *res = append(*res, string(*candidate))
        return
    }
    
    curNum := int(digits[cur] - '0')
    for i := 0; i < len(letterMap[curNum]); i++ {
        char := letterMap[curNum][i]
        *candidate = append(*candidate, char)
        helper(digits, cur + 1, candidate, res, letterMap)
        *candidate = (*candidate)[:len(*candidate) - 1]
    }
}
