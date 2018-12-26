func ladderLength(beginWord string, endWord string, wordList []string) int {
    dict := make(map[string]struct{})
    for _, word := range wordList {
        dict[word] = struct{}{}
    }
    
    // BFS
    res := 1 
    queue := make([]string, 0)
    queue = append(queue, beginWord)
    visited := make(map[string]struct{})
    for len(queue) > 0 {
        size := len(queue)
        res++
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            next := getNextWords(cur, dict)
            for _, word := range next {
                _, ok := visited[word]
                if ok {
                    continue
                }
                visited[word] = struct{}{}
                queue = append(queue, word)
                if word == endWord {
                    return res
                }
            }
        }
    }
    return 0
}

func getNextWords(cur string, dict map[string]struct{}) []string {
    res := make([]string, 0)
    curBytes := []byte(cur)
    for i := 0; i < len(cur); i++ {
        var c byte = 'a'
        for c = 'a'; c <= 'z'; c++ {
            curBytes[i] = c
            candidate := string(curBytes)
            _, ok := dict[candidate]
            if ok {
                res = append(res, candidate)
            }
            curBytes[i] = cur[i]
        }
    }
    return res
}
