func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    cache := make([][]int, 0)
    for i := 0; i <= m; i++ {
        cache = append(cache, make([]int, n + 1))
    }
    res := helper(s, p, 0, 0, cache)
    if res == 1 {
        return true
    } else {
        return false
    }
}

func helper(s string, p string, startS int, startP int, cache [][]int) int {
    if cache[startS][startP] != 0 {
        return cache[startS][startP]
    }
    if startS == len(s) && startP == len(p) {
        cache[startS][startP] = 1
        return 1
    }
    
    if startP == len(p) {
        cache[startS][startP] = -1
        return -1
    }
    
    res := -1
    // Normal case
    if startP + 1 >= len(p) || p[startP + 1] != '*' {
        if startS >= len(s) {
            res = -1
        } else if p[startP] == '.' || p[startP] == s[startS] {
            res = helper(s, p, startS + 1, startP + 1, cache)
        } else {
            res = -1
        }
        cache[startS][startP] = res
        return res
    }   
    
    // Case of *
    char := p[startP]
    curS := startS
    if helper(s, p, curS, startP + 2, cache) == 1 {
        cache[startS][startP] = 1
        return 1
    }
    for curS < len(s) && (char == '.' || char == s[curS]) {
        curS++
        if helper(s, p, curS, startP + 2, cache) == 1 {
            cache[startS][startP] = 1
            return 1
        }
    }
    cache[startS][startP] = -1
    return -1
}

