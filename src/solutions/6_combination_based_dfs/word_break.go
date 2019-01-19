func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]struct{})
    for _, word := range wordDict {
        dict[word] = struct{}{}
    }
    
    res := make([]bool, len(s) + 1)
    res[0] = true
    for i := 0; i < len(s); i++ {
        if !res[i] {
            continue
        }
        for j := i; j < len(s); j++ {
            _, ok := dict[s[i : j + 1]]
            if ok {
                res[j + 1] = true
            }
        }
    }
    return res[len(s)]
}
