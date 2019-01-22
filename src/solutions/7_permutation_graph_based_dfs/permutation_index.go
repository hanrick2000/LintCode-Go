/**
 * @param A: An array of integers
 * @return: A long integer
 */
func permutationIndex (A []int) int64 {
    if len(A) == 0 {
        return 0
    }
    n := len(A)
    
    factor := factorial(n)
    var res int64 = 0
    for i := 0; i < n; i++ {
        cur := A[i]
        smallerNum := 0
        for j := i + 1; j < n; j++ {
            if A[j] < cur {
                smallerNum++
            }
        }
        res += int64(factor[n - 1 - i] * smallerNum)
    }
    return res + 1
}

func factorial(n int) []int {
    res := make([]int, 0)
    res = append(res, 1)
    cur := 1
    for i := 1; i < n; i++ {
        cur *= i
        res = append(res, cur) 
    }
    return res
}
