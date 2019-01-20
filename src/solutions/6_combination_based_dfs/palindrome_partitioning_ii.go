func minCut(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    isMatrix := make([][]bool, 0)
    for i := 0; i < n; i++ {
        isMatrix = append(isMatrix, make([]bool, n))
    }
    createMatrix(s, &isMatrix)
    
    res := make([]int, n + 1)
    res[0] = -1
    for i := 1; i <= n; i++ {
        res[i] = res[i - 1] + 1
        for j := 0; j < i; j++ {
            if isMatrix[j][i - 1] && res[j] + 1 < res[i] {
                res[i] = res[j] + 1
            }
        }
    }
    return res[n]
}

func createMatrix(s string, isMatrix *[][]bool) {
    onePass(s, 0, 0, isMatrix)
    for i := 1; i < len(s); i++ {
        onePass(s, i, i, isMatrix)
        onePass(s, i - 1, i, isMatrix)
    }
}

func onePass(s string, left int, right int, isMatrix *[][]bool) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        (*isMatrix)[left][right] = true
        left--
        right++
    }
}
