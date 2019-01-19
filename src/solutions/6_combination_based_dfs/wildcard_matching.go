func isMatch(s string, p string) bool {
    curS, curP := 0, 0
    prevS, prevP := -1, -1
    
    for curS < len(s) || curP < len(p) {
        if curP == len(p) {
            if prevP != -1 && prevS < len(s){
                prevS++
                curP = prevP
                curS = prevS
                continue
            } else {
                return false
            }
        }
        
        if curP < len(p) && p[curP] == '*' {
                curP++
                prevP = curP
                prevS = curS
                continue
        }
        
        if curS == len(s) {
            return false
        } 
        
        if p[curP] == '?' || p[curP] == s[curS] {
            curP++
            curS++
        } else if prevP != -1 && prevS < len(s) {
            prevS++
            curP = prevP
            curS = prevS
        } else {
            return false
        }
    }
    return true
}
