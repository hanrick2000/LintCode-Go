/**
 * @param image: a binary matrix with '0' and '1'
 * @param x: the location of one of the black pixels
 * @param y: the location of one of the black pixels
 * @return: an integer
 */
 
func search(image [][]byte, leftx int, lefty int, rightx int, righty int, leftmost bool) [2]int {
    for leftx + 1 < rightx || lefty + 1 < righty {
        midx := leftx + (rightx - leftx) / 2
        midy := lefty + (righty - lefty) / 2
        if image[midx][midy] == '1' {
            if leftmost {
                rightx, righty = midx, midy
            } else {
                leftx, lefty = midx, midy
            }
        } else {
            if leftmost {
                leftx, lefty = midx, midy
            } else {
                rightx, righty = midx, midy
            }
        }
    }
    if leftmost {
        if image[leftx][lefty] == '1' {
            return [2]int{leftx, lefty}
        } else {
            return [2]int{rightx, righty}
        }
    } else {
        if image[rightx][righty] == '1' {
            return [2]int{rightx, righty}
        } else {
            return [2]int{leftx, lefty}
        }
    }
}

func update(res *[4]int, x int, y int, larger bool) {
    if larger {
        if (*res)[1] < x {
            (*res)[1] = x
        }
        if (*res)[3] < y {
            (*res)[3] = y
        }
    } else {
        if (*res)[0] > x {
            (*res)[0] = x
        }
        if (*res)[2] > y {
            (*res)[2] = y
        }
    }
}

// Solution 1: DFS version
func helperDFS(image [][]byte, x int, y int, res *[4]int, visited map[int]struct{}) {
    m, n := len(image), len(image[0])
    _, ok := visited[x * n + y]
    if ok {
        return
    }
    visited[x * n + y] = struct{}{}
    next := search(image, x, 0, x, y, true)
    update(res, next[0], next[1], false)
    helperDFS(image, next[0], next[1], res, visited)
    
    next = search(image, x, y, x, n - 1, false)
    update(res, next[0], next[1], true)
    helperDFS(image, next[0], next[1], res, visited)
    
    next = search(image, 0, y, x, y, true)
    update(res, next[0], next[1], false)
    helperDFS(image, next[0], next[1], res, visited)
    
    next = search(image, x, y, m - 1, y, false)
    update(res, next[0], next[1], true)
    helperDFS(image, next[0], next[1], res, visited)
}

// Solution 2: BFS version
func helperBFS(image [][]byte, x int, y int, res *[4]int, visited map[int]struct{}) {
    m, n := len(image), len(image[0])
    queue := make([]int, 0)
    queue = append(queue, x * n + y)
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        _, ok := visited[cur]
        if ok {
            continue
        }
        visited[cur] = struct{}{}
        x = cur / n
        y = cur % n
        
        next := search(image, x, 0, x, y, true)
        update(res, next[0], next[1], false)
        queue = append(queue, next[0] * n + next[1])
    
        next = search(image, x, y, x, n - 1, false)
        update(res, next[0], next[1], true)
        queue = append(queue, next[0] * n + next[1])
    
        next = search(image, 0, y, x, y, true)
        update(res, next[0], next[1], false)
        queue = append(queue, next[0] * n + next[1])
    
        next = search(image, x, y, m - 1, y, false)
        update(res, next[0], next[1], true)
        queue = append(queue, next[0] * n + next[1])
    }
}
 
func minArea (image [][]byte, x int, y int) int {
    visited := make(map[int]struct{})
    res := [4]int{x, x, y, y}
    helperBFS(image, x, y, &res, visited)
    return (res[1] - res[0] + 1) * (res[3] - res[2] + 1)
}

// 0 0 0 2
// 1 2
