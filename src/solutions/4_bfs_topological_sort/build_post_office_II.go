/**
 * @param grid: a 2D grid
 * @return: An integer
 */
func shortestDistance (grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    distanceMatrix := make([][][]int, 0)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                distanceMatrix = append(distanceMatrix, getDistanceMatrix(i, j, grid))
            }
        }
    }
    
    // Get final res
    res := -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                continue
            }
            cur := 0
            reachable := true
            for k := 0; k < len(distanceMatrix); k++ {
                if distanceMatrix[k][i][j] == 0 {
                    reachable = false
                    break
                }
                cur += distanceMatrix[k][i][j]
            }
            if reachable && (res == -1 || cur < res) {
                res = cur
            }
        }
    }
    return res
}

func getDistanceMatrix (startx int, starty int, grid [][]int) [][]int {
    start := [2]int{startx, starty}
    queue := make([][2]int, 0)
    queue = append(queue, start)
    visited := make(map[[2]int]struct{})
    visited[start] = struct{}{}
    m, n := len(grid), len(grid[0])
    level := 0
    res := make([][]int, 0)
    for i := 0; i < m; i++ {
        res = append(res, make([]int, n))
    }
    
    for len(queue) > 0 {
        size := len(queue)
        level++
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            nextx := []int{-1, 1, 0, 0}
            nexty := []int{0, 0, -1, 1}
            for j := 0; j < 4; j++ {
                if isValid(cur[0] + nextx[j], cur[1] + nexty[j], grid, visited) {
                    next := [2]int{cur[0] + nextx[j], cur[1] + nexty[j]}
                    queue = append(queue, next)
                    visited[next] = struct{}{}
                    res[next[0]][next[1]] = level
                }
            }
        }
    }
    return res
}

func isValid(x int, y int, grid [][]int, visited map[[2]int]struct{}) bool {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
        return false
    }
    if grid[x][y] != 0 {
        return false
    }
    _, ok := visited[[2]int{x, y}]
    if ok {
        return false
    }
    return true
}

