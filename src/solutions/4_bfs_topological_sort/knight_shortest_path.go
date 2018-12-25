/**
 * Definition for a point.
 * type Point struct {
 *     X, Y int
 * }
 */

/**
 * @param grid: a chessboard included 0 (false) and 1 (true)
 * @param source: a point
 * @param destination: a point
 * @return: the shortest path 
 */
func shortestPath (grid [][]bool, source *Point, destination *Point) int {
    res := 0
    queue := make([]Point, 0)
    visited := make(map[Point]struct{})
    nextx := []int{1, 1, -1, -1, 2, 2, -2, -2}
    nexty := []int{2, -2, 2, -2, 1, -1, 1, -1}
    queue = append(queue, *source)
    m := len(grid)
    n := len(grid[0])
    
    for len(queue) > 0 {
        res++
        size := len(queue)
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            for j := 0; j < 8; j++ {
                next := Point{cur.X + nextx[j], cur.Y + nexty[j]}
                if next == *destination {
                    return res
                }
                if valid(next, m, n, visited, grid) {
                    visited[next] = struct{}{}
                    queue = append(queue, next)
                }
            }
        }
    }
    return -1
}

func valid(next Point, m int, n int, visited map[Point]struct{}, grid [][]bool) bool {
    if next.X < 0 || next.X >= m || next.Y < 0 || next.Y >= n {
        return false
    }
    _, ok := visited[next]
    if ok {
        return false
    }
    if grid[next.X][next.Y] {
        return false
    }
    return true
}

