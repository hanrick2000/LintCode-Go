func wordPattern(pattern string, str string) bool {
    words := strings.Split(str, " ")
    wordMap := make(map[byte]string)
    charMap := make(map[string]byte)

    if len(words) != len(pattern) {
        return false
    }
    
    for i := 0; i < len(pattern); i++ {
        if word, ok := wordMap[pattern[i]]; ok {
            if word != words[i] {
                return false
            }
        } else {
            wordMap[pattern[i]] = words[i]
            if char, ok := charMap[words[i]]; ok {
                if char != pattern[i] {
                    return false
                }
            }
            charMap[words[i]] = pattern[i]
        }
    }
    return true
}
