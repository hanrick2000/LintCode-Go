func wordPatternMatch(pattern string, str string) bool {
    char2Word := make(map[byte]string)
    word2Char := make(map[string]byte)
    
    return helper(pattern, str, 0, 0, char2Word, word2Char)
}

func helper(pattern string, str string, startP int, startS int, 
            char2Word map[byte]string, word2Char map[string]byte) bool {
    if startP == len(pattern) && startS == len(str) {
        return true
    }
    if startP == len(pattern) || startS == len(str) {
        return false
    }
    
    if word, ok := char2Word[pattern[startP]]; ok {
        for i := startS; i < len(str); i++ {
            if str[startS:i + 1] == word {
                return helper(pattern, str, startP + 1, i + 1, char2Word, word2Char)
            }
        }
    } else {
        // the case that we don't have any word for the current char
        for i := startS; i < len(str); i++ {
            word = str[startS:i + 1]
            if _, ok := word2Char[word]; ok {
                continue
            }   
            char2Word[pattern[startP]] = word
            word2Char[word] = pattern[startP]
            if helper(pattern, str, startP + 1, i + 1, char2Word, word2Char) {
                return true
            }
            delete(char2Word, pattern[startP])
            delete(word2Char, word)
        }
    }
    return false
}
