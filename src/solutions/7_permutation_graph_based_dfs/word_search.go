func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return word == ""
    }
    
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if findWord(board, i, j, word, 0) {
                return true
            }
        }
    }
    return false
}

func findWord(board [][]byte, x int, y int, 
              word string, cur int) bool {
    m, n := len(board), len(board[0])
    if cur >= len(word) {
        return true
    }
    
    if board[x][y] != word[cur] {
        return false
    }
    
    if cur + 1 == len(word) {
        return true
    }
    
    temp := board[x][y]
    board[x][y] = '*'
    deltas := [][]int {
        {0, 1}, {0, -1},
        {1, 0}, {-1, 0},
    }
    for _, delta := range deltas {
        next := [2]int{x + delta[0], y + delta[1]}
        if next[0] < 0 || next[0] >= m || next[1] < 0 || 
            next[1] >= n || board[next[0]][next[1]] == '*' {
            continue
        }
        if findWord(board, next[0], next[1], word, cur + 1) {
            return true
        }
    }
    board[x][y] = temp
    return false
}
