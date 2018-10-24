// KMP algorithm

func strStr(haystack string, needle string) int {
    n := len(needle)
    if n == 0 {
        return 0
    }
    kmpTable := getKMPTable(needle)
    return searchKMP(haystack, needle, kmpTable)
}

func getKMPTable(needle string) []int {
    res := make([]int, len(needle))
    res[0] = -1
    cur := 1
    prev := 0
    for ; cur < len(needle); {
        if needle[cur] == needle[prev] {
            res[cur] = prev
            cur++
            prev++
        } else if prev == 0 {
            res[cur] = -1
            cur++
            prev = 0
        } else {
            prev = res[prev - 1] + 1
        }
    }
    fmt.Println(res)
    return res
}

func searchKMP(haystack string, needle string, kmpTable []int) int {
    m, n := len(haystack), len(needle)
    cur := 0
    matched := 0
    for ; cur < m; {
        if cur + n > m {
            break
        }
        if haystack[cur + matched] == needle[matched] {
            matched++
            if matched == n {
                return cur
            }
        } else if matched == 0 {
            cur++
        } else {
            end := cur + matched
            matched = kmpTable[matched - 1] + 1
            cur = end - matched
        }
    }
    return -1
}
