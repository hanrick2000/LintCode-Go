func strStr(haystack string, needle string) int {
    m, n := len(haystack), len(needle)
    if n == 0 {
        return 0
    }
    for i := 0; i < m; i++ {
        if i + n <= m {
            cur := 0
            for ;cur < n; cur++ {
                if haystack[i + cur] != needle[cur] {
                    break
                }
            }
            if cur == n {
                return i
            }
        }
    }
    return -1
}
