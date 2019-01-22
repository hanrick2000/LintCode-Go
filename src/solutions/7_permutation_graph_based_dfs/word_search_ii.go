type trie struct {
    root *trieNode
}

type trieNode struct {
    nextNode map[byte]*trieNode
}

func newTrie() trie {
    root := trieNode{make(map[byte]*trieNode)}
    return trie{&root}
}

func (curTrie *trie) insert(word string) {
    cur := curTrie.root
    for i := 0; i < len(word); i++ {
        if next, ok := cur.nextNode[word[i]]; ok {
            cur = next
        } else {
            cur.nextNode[word[i]] = &trieNode{make(map[byte]*trieNode)}
            cur = cur.nextNode[word[i]]
        }
    }
    cur.nextNode['$'] = nil
}

func findWords(board [][]byte, words []string) []string {
    res := make(map[string]struct{})
    if len(board) == 0 {
        return make([]string, 0) 
    }
    m, n := len(board), len(board[0])
    
    // Create trie
    wordTrie := newTrie()
    for _, word := range words {
        wordTrie.insert(word)
    }
    
    // Run dfs
    candidate := make([]byte, 0)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(board, wordTrie.root, i, j, &candidate, res)
        }
    }
    
    finalRes := make([]string, 0)
    for k, _ := range res {
        finalRes = append(finalRes, k)
    }
    return finalRes
}

func dfs(board [][]byte, cur *trieNode, x int, y int,
         candidate *[]byte, res map[string]struct{}) {
    m, n := len(board), len(board[0])
    char := board[x][y]

    if _, ok := cur.nextNode[char]; !ok {
        return
    }
    
    cur = cur.nextNode[char]
    *candidate = append(*candidate, char)
    board[x][y] = '*'
    if _, ok := cur.nextNode['$']; ok {
        res[string(*candidate)] = struct{}{}
    }
    
    deltas := [][]int{
        {0, 1}, {0, -1},
        {-1, 0}, {1, 0},
    }
    for _, delta := range deltas {
        nextx := x + delta[0]
        nexty := y + delta[1]
        if nextx < 0 || nextx >= m || nexty < 0 ||
           nexty >= n || board[nextx][nexty] == '*' {
           continue 
        } 
        dfs(board, cur, nextx, nexty, candidate, res)
    }
    *candidate = (*candidate)[:len(*candidate) - 1]
    board[x][y] = char 
}

