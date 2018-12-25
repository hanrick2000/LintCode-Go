func canFinish(numCourses int, prerequisites [][]int) bool {
    if numCourses == 0 {
        return true
    }
    graph := buildGraph(numCourses, prerequisites)
    inDegree := countIndegree(numCourses, graph)
    zeroDegree := getZeroDegree(inDegree)
    res := 0
    
    for len(zeroDegree) > 0 {
        res++
        n := len(zeroDegree)
        nextTake := zeroDegree[n - 1]
        zeroDegree = zeroDegree[0 : n - 1]
        for _, val := range graph[nextTake] {
            inDegree[val]--
            if inDegree[val] == 0 {
                zeroDegree = append(zeroDegree, val)
            }
        }
    }
    return res == numCourses
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
    graph := make([][]int, numCourses)
    for _, prev := range prerequisites {
        graph[prev[1]] = append(graph[prev[1]], prev[0])
    }
    return graph
}

func countIndegree(n int, graph [][]int) []int {
    res := make([]int, n)
    for _, nextNodes := range graph {
        for _, next := range nextNodes {
            res[next]++
        }
    }
    return res
}

func getZeroDegree(inDegree []int) []int {
    res := make([]int, 0)
    for i, val := range inDegree {
        if val == 0 {
            res = append(res, i)
        }
    }
    return res
}
