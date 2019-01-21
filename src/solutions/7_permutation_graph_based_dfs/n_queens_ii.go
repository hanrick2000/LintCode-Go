func totalNQueens(n int) int {
    bitmap := make([][]int, 0)
    for i := 0; i < n; i++ {
        bitmap = append(bitmap, make([]int, n))
    }
    res := 0
    helper(0, n, bitmap, &res)
    return res
}

func helper(cur int, n int, bitmap [][]int, res *int) {
    if cur == n {
        *res = *res + 1
        return
    }
    
    for i := 0; i < n; i++ {
        if valid(bitmap, cur, i) {
            bitmap[cur][i] = 1
            helper(cur + 1, n, bitmap, res)
            bitmap[cur][i] = 0
        }
    }
}

func valid(bitmap [][]int, x int, y int) bool {
    n := len(bitmap)
    for i := 0; i < n; i++ {
        if bitmap[x][i] == 1 || bitmap[i][y] == 1 {
            return false
        }
    }
    
    toAdd := [][]int{
        {0, 1}, {0, -1},
        {1, 0}, {-1, 0},
        {1, 1}, {1, -1},
        {-1, 1}, {-1, -1},
    }
    
    for _, add := range toAdd {
        cur := []int{x, y}
        for cur[0] >= 0 && cur[0] < n && cur[1] >= 0 && cur[1] < n {
            if bitmap[cur[0]][cur[1]] == 1 {
                return false
            }
            cur[0] += add[0]
            cur[1] += add[1]
        } 
    }
    return true
}

