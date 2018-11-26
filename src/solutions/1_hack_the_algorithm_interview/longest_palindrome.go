func longestPalindrome (s string) string {
    res := ""
    for i := 0; i < len(s); i++ {
        candidate := helper(s, i, i)
        if len(candidate) > len(res) {
            res = candidate
        }
        if i != len(s) - 1 {
            candidate = helper(s, i, i+1)
            if len(candidate) > len(res) {
                res = candidate
            }   
        }
    }
    return res
}

func helper(s string, left int, right int) (res string) {
    for right < len(s) && left >= 0 && s[left] == s[right] {
        left -= 1
        right += 1
    }
    return s[left+1 : right]
}
