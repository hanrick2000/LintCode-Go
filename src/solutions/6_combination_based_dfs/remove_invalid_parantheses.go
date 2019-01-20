func removeInvalidParentheses(s string) []string {
    visited := make(map[string]struct{})
    removedNum := getRemovedNum(s)
    candidate := make([]byte, 0)
    res := make([]string, 0)
    helper(s, 0, 0, 0, 0, removedNum, &candidate, &res, visited)
    return res
}

func helper(s string, cur int, left int, right int, 
            curRemoved int, targetRemoved int, 
            candidate *[]byte, res *[]string, visited map[string]struct{}) {
    if cur == len(s) {
        if left == right && curRemoved == targetRemoved {
            tempString := string(*candidate)
            if _, ok := visited[tempString]; !ok {
                *res = append(*res, string(*candidate))
                visited[tempString] = struct{}{}
            }
        }
        return
    }
    
    // Remove current character
    if (s[cur] == '(' || s[cur] == ')') && curRemoved < targetRemoved {
        helper(s, cur + 1, left, right, 
               curRemoved + 1, targetRemoved, candidate, res, visited)
    }
    
    // Add if '(' or ')'
    if s[cur] == '(' {
        left++
    } else if s[cur] == ')' {
        right++
    }
    if left >= right {
        *candidate = append(*candidate, s[cur])
        helper(s, cur + 1, left, right, 
               curRemoved, targetRemoved, candidate, res, visited)
        *candidate = (*candidate)[:len(*candidate) - 1]
    }
}

func getRemovedNum(s string) int {
    res := 0
    cur := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            cur++
        } else if s[i] == ')' {
            cur--
        }
        if cur < 0 {
            res++
            cur = 0
        }
    }
    cur = 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ')' {
            cur++
        } else if s[i] == '(' {
            cur--
        }
        if cur < 0 {
            res++
            cur = 0
        }
    }
    return res
}
