func slidingPuzzle(board [][]int) int {
    goal := [6]int{1, 2, 3, 4, 5, 0}
    visited := make(map[[6]int]struct{})
    queue := make([][6]int, 0)
    queue = append(queue, getStatus(board))
    visited[getStatus(board)] = struct{}{}
    level := -1
    
    for len(queue) > 0 {
        size := len(queue)
        level++
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            if cur == goal {
                return level
            }
            for _, next := range getNextMoves(cur) {
                _, ok := visited[next]
                if !ok {
                    visited[next] = struct{}{}
                    queue = append(queue, next)
                }
            }
        }
    }
    return -1
}

func getStatus(board [][]int) [6]int {
    return [6]int{
        board[0][0], board[0][1], board[0][2],
        board[1][0], board[1][1], board[1][2],
    }
}

func getNextMoves(cur [6]int) [][6]int {
    res := make([][6]int, 0)
    m, n := 0, 0
    for i := 0; i < 6; i++ {
        if cur[i] == 0 {
            m = i / 3
            n = i % 3
        }
    }
    // Change row
    next := cur
    next[n], next[3 + n] = next[3 + n], next[n]
    res = append(res, next)
   
    // Change column
    if n < 2 {
        next = cur
        next[3 * m + n], next[3 * m + n + 1] = next[3 * m + n + 1], next[3 * m + n]
        res = append(res, next)
    }
    if n > 0 {
        next = cur
        next[3 * m + n], next[3 * m + n - 1] = next[3 * m + n - 1], next[3 * m + n]
        res = append(res, next)
    }
    return res
}

