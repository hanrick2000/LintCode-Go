func countComponents(n int, edges [][]int) int {
    visited := make(map[int]struct{})
    graph := make([][]int, n)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    res := 0
    for i := 0; i < n; i++ {
        _, ok := visited[i]
        if !ok {
            res++
            dfs(i, visited, graph)
        }
    }
    return res
}

func dfs(cur int, visited map[int]struct{}, graph [][]int) {
    _, ok := visited[cur]
    if ok {
        return
    }
    visited[cur] = struct{}{}
    for _, next := range graph[cur] {
        dfs(next, visited, graph)
    }
}
