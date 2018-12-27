/**
 * @param grid: a 2D grid
 * @return: An integer
 */
func shortestDistance (grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    hdistance := make([]int, n)
    vdistance := make([]int, m)
    xs, ys := make([]int, 0), make([]int, 0)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                xs = append(xs, i)
                ys = append(ys, j)
            }
        }
    }
    
    for i := 0; i < n; i++ {
        for _, x := range ys {
            hdistance[i] += abs(x - i)
        }
    }
    for j := 0; j < m; j++ {
        for _, y := range xs {
            vdistance[j] += abs(y - j)
        }
    }
    
    res := -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                cur := vdistance[i] + hdistance[j]
                if res == -1 || cur < res {
                    res = cur
                }
            }
        }
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

