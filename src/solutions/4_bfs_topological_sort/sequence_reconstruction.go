
func sequenceReconstruction(org []int, seqs [][]int) bool {
    graph := buildGraph(seqs)
    n := len(graph)
    if len(org) != n {
        return false
    }
    inDegree := getInDegree(seqs)
    zeroDegree := getZeroDegree(inDegree)
    res := make([]int, 0)
    
    for len(zeroDegree) != 0 {
        if len(zeroDegree) != 1 {
            return false
        }
        nextNode := zeroDegree[0]
        zeroDegree = zeroDegree[1:]
        res = append(res, nextNode)
        for _, val := range graph[nextNode] {
            inDegree[val]--
            if inDegree[val] == 0 {
                zeroDegree = append(zeroDegree, val)
            }
        }
    }
    if len(res) != len(org) {
        return false
    }
    for i := 0; i < len(res); i++ {
        if res[i] != org[i] {
            return false
        }
    }
    return true
}

func buildGraph(seqs [][]int) map[int][]int {
    res := make(map[int][]int)
    for _, seq := range seqs {
        for _, val := range seq {
            _, ok := res[val]
            if !ok {
                res[val] = []int{}
            }
        }
    }
    for _, seq := range seqs {
        for i := 0; i < len(seq) - 1; i++ {
            nextNodes, _ := res[seq[i]]
            res[seq[i]] = append(nextNodes, seq[i + 1])
        }
    }
    return res
}

func getInDegree(seqs [][]int) map[int]int {
    res := make(map[int]int)
    for _, seq := range seqs {
        for _, val := range seq {
            _, ok := res[val]
            if !ok {
                res[val] = 0
            }
        }
    }
    for _, seq := range seqs {
       for i := 0; i < len(seq) - 1; i++ {
            res[seq[i + 1]]++
       }
    }
    return res
}

func getZeroDegree(inDegree map[int]int) []int {
    res := make([]int, 0)
    for key, value := range inDegree {
        if value == 0 {
            res = append(res, key)
        }
    }
    return res
}
