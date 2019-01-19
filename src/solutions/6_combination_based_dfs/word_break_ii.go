func wordBreak(s string, wordDict []string) []string {
    dict := make(map[string]struct{})
    for _, word := range wordDict {
        dict[word] = struct{}{}
    }
    
    n := len(s)
    res := make(map[int][][]byte)
    res[0] = make([][]byte, 0)
    res[0] = append(res[0], make([]byte, 0))
    
    dfs(n, s, dict, res)
    
    finalRes := make([]string, 0)
    for _, word := range res[n] {
        finalRes = append(finalRes, string(word))
    }
    return finalRes
}

func dfs(cur int, s string, dict map[string]struct{}, res map[int][][]byte) [][]byte {
    if val, ok := res[cur]; ok {
        return val 
    }
    curRes := make([][]byte, 0)
    for i := 0; i < cur; i++ {
        curWord := s[i:cur]
        if _, ok := dict[curWord]; ok {
            for _, words := range dfs(i, s, dict, res) {
                temp := make([]byte, 0)
                temp = append(temp, words...)
                if len(temp) > 0 {
                    temp = append(temp, ' ')
                }
                temp = append(temp, []byte(curWord)...)
                curRes = append(curRes, temp)
            }
        }
    }
    
    res[cur] = curRes
    return curRes
}
