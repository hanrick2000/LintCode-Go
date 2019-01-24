type trie struct {
    root *trieNode
}

type trieNode struct {
    next map[byte]*trieNode
}

func (wordTrie *trie) insert(word string) {
    cur := wordTrie.root
    for i := 0; i < len(word); i++ {
        if _, ok := cur.next[word[i]]; !ok {
            newNode := trieNode{make(map[byte]*trieNode)}
            cur.next[word[i]] = &newNode
        }
        cur = cur.next[word[i]]
    }
    cur.next['$'] = nil
}

func (wordTrie *trie) findAll(prefix string, cur *trieNode, res *[]string) {
    for char, nextNode := range cur.next {
        if char == '$' {
            *res = append(*res, prefix)
        } else {
            wordTrie.findAll(prefix + string(char), nextNode, res)
        }
    }
}

func (wordTrie *trie) getWords(prefix string) []string {
    cur := wordTrie.root
    res := make([]string, 0)
    for i := 0; i < len(prefix); i++ {
        if _, ok := cur.next[prefix[i]]; !ok {
            return res
        }
        cur = cur.next[prefix[i]]
    }
    wordTrie.findAll(prefix, cur, &res)
    return res
}

func newTrie() trie {
    root := trieNode{make(map[byte]*trieNode)}
    return trie{&root}
}

func wordSquares(words []string) [][]string {
    wordTrie := newTrie()
    for _, word := range words {
        wordTrie.insert(word)
    }
    
    res := make([][]string, 0)
    candidate := make([]string, 0)
    for _, word := range words {
        n := len(word)
        candidate = append(candidate, word)
        dfs(1, n, &candidate, &res, wordTrie)
        candidate = candidate[0 : len(candidate) - 1]
    }
    return res
}

func dfs(cur int, n int, candidate *[]string, res *[][]string, wordTrie trie) {
    if cur == n {
        toAdd := make([]string, len(*candidate))
        copy(toAdd, *candidate)
        *res = append(*res, toAdd)
        return
    }
    
    // Now search for words
    prefix := ""
    for _, word := range *candidate {
        prefix += string(word[cur])
    }
    for _, word := range wordTrie.getWords(prefix) {
        *candidate = append(*candidate, word)
        dfs(cur + 1, n, candidate, res, wordTrie)
        *candidate = (*candidate)[0:len(*candidate) - 1]
    }
}

