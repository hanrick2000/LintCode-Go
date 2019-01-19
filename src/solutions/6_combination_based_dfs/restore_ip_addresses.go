func restoreIpAddresses(s string) []string {
    res := make([]string, 0)
    candidate := make([]string, 0)
    helper(s, 0, &candidate, &res)
    return res
}

func helper(s string, start int, candidate *[]string, res *[]string) {
    if start == len(s) && len(*candidate) == 4 {
        toAdd := ""
        for _, num := range *candidate {
            toAdd += num
            toAdd += "."
        }
        toAdd = toAdd[0 : len(toAdd) - 1]
        *res = append(*res, toAdd)
        return
    }
    
    if start == len(s) || len(*candidate) == 4 {
        return
    }
    
    for i := 1; i <= 3; i++ {
        if start + i <= len(s) && isValid(s[start : start + i]) {
            *candidate = append(*candidate, s[start : start + i])
            helper(s, start + i, candidate, res)
            *candidate = (*candidate)[0 : len(*candidate) - 1]
        }
    }
}

func isValid(s string) bool {
    if len(s) == 1 {
        return true
    }
    if s[0] == '0' {
        return false
    }
    if len(s) == 2 {
        return true
    }
    num := int(s[0] - '0') * 100 + int(s[1] - '0') * 10 + int(s[2] - '0')
    if num > 255 {
        return false
    }
    return true
}
