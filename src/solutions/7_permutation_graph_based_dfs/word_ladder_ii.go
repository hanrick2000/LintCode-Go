type graphNode struct {
    level int
    str string
    next []*graphNode
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    wordSet := make(map[string]struct{})
    nodeMap := make(map[string]*graphNode)
    queue := make([]string, 0)
    
    // queue initialization
    queue = append(queue, beginWord)
    curNode := graphNode{0, beginWord, make([]*graphNode, 0)}
    nodeMap[beginWord] = &curNode
    level := 0
    found := false
    
    for _, word := range wordList {
        wordSet[word] = struct{}{}
    }
    
    for len(queue) > 0 && !found {
        size := len(queue)
        level++
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            for _, next := range nextWords(cur, wordSet) {
                if nextNode, ok := nodeMap[next]; !ok {
                    // The case the next word not visited
                    queue = append(queue, next)
                    addNode := graphNode{level, next, make([]*graphNode, 0)}
                    addNode.next = append(addNode.next, nodeMap[cur])
                    nodeMap[next] = &addNode
                } else {
                    // The case next word already visited, need to check
                    if nextNode.level != level {
                        continue
                    }
                    nextNode.next = append(nextNode.next, nodeMap[cur])
                }
                
                if next == endWord {
                    found = true
                }
            }
        }
    }
    
    // DFS the graphNode for final res
    res := make([][]string, 0)
    if endNode, ok := nodeMap[endWord]; !ok {
        return res
    } else {
        candidate := make([]string, 0)
        DFS(endNode, beginWord, &candidate, &res)
    }
    
    // Convert to final results, reverse the order
    for _, oneResult := range res {
        for i := 0; i < len(oneResult) / 2; i++ {
            oneResult[i], oneResult[len(oneResult) - i - 1] = 
                oneResult[len(oneResult) - i - 1], oneResult[i]
        }
    }
    return res
}

func DFS(curNode *graphNode, targetWord string, 
         candidate *[]string, res *[][]string) {
    
    *candidate = append(*candidate, curNode.str)
    if curNode.str == targetWord {
        toAdd := make([]string, len(*candidate))
        copy(toAdd, *candidate)
        *res = append(*res, toAdd)
    }
    for _, nextNode := range curNode.next {
        DFS(nextNode, targetWord, candidate, res)
    }
    *candidate = (*candidate)[:len(*candidate) - 1]
}

func nextWords(word string, wordSet map[string]struct{}) []string {
    res := make([]string, 0)
    wordBytes := []byte(word)
    var char byte
    
    for i := 0; i < len(wordBytes); i++ {
        for char = 'a'; char <= 'z'; char++ {
            if char == word[i] {
                continue
            }
            wordBytes[i] = char
            curWord := string(wordBytes)
            if _, ok := wordSet[curWord]; ok {
                res = append(res, curWord)
            }
            wordBytes[i] = word[i]
        }
    }
    return res
}
