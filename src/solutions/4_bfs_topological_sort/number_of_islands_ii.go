func numIslands2(m int, n int, positions [][]int) []int {
    res := make([]int, len(positions))
    if m == 0 || n == 0 {
        return res
    }
    disjointSet := make([]int, m * n)
    for i := 0; i <  len(disjointSet); i++ {
            disjointSet[i] = -1
    }
    
    count := 0
    for i, position := range positions {
        nextx := []int{-1, 1, 0, 0}
        nexty := []int{0, 0, -1, 1}
        disjointSet[position[0] * n + position[1]] = position[0] * n + position[1]
        count++
        for i := 0; i < 4; i++ {
            curx := position[0] + nextx[i]
            cury := position[1] + nexty[i]
            if curx >= 0 && curx < m && cury >= 0 && cury < n {
                root1 := find(disjointSet, position[0] * n + position[1])
                root2 := find(disjointSet, curx * n + cury)
                if root2 == -1 {
                    continue
                }
                if root1 == root2 {
                    continue
                }
                disjointSet[root1] = root2
                count--
            }
        }
        res[i] = count
    }
    return res
}

func find(disjointSet []int, cur int) int {
    if disjointSet[cur] == cur {
        return cur
    }
    if disjointSet[cur] == -1 {
        return -1
    }
    return find(disjointSet, disjointSet[cur])
}
