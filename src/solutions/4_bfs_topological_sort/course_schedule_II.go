func findOrder(numCourses int, prerequisites [][]int) []int {
    res := make([]int, 0)
    if numCourses == 0 {
        return res
    }
    
    graph := buildGraph(numCourses, prerequisites)
    visited := make(map[int]int)
    for i := 0; i < numCourses; i++ {
        visited[i] = 0
    }
    for i := 0; i < numCourses; i++ {
        if visited[i] != 0 {
            continue
        }
        if !dfs(graph, visited, i, &res) {
            return []int{}
        }
    }
    
    // reverse res
    for i := 0; i < len(res) / 2; i++ {
        res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
    }
    return res
}

func dfs(graph [][]int, visited map[int]int, start int, res *[]int) bool {
    visited[start] = 2
    for _, val := range graph[start] {
        // 2 means visited in the current cycle
        if visited[val] == 2 {
            return false
        } else if visited[val] == 1 {
            continue
        } else if !dfs(graph, visited, val, res) {
            return false
        }
    }
    visited[start] = 1
    *res = append(*res, start)
    return true
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
    graph := make([][]int, numCourses)
    for _, prev := range prerequisites {
        graph[prev[1]] = append(graph[prev[1]], prev[0])
    }
    return graph
}
