func mergekSortedArrays (arrays [][]int) []int {
    n := len(arrays)
    return helper(arrays, 0, n)
}

func helper(arrays [][]int, start int, end int) []int {
    if start == end {
        return make([]int, 0)
    }
    if end - start == 1 {
        return arrays[start]
    }
    left := helper(arrays, start, (start + end) / 2)
    right := helper(arrays, (start + end) / 2, end)
    res := make([]int, 0)
    leftCur, rightCur := 0, 0
    for leftCur < len(left) || rightCur < len(right) {
        if leftCur == len(left) {
            res = append(res, right[rightCur])
            rightCur++
        } else if rightCur == len(right) {
            res = append(res, left[leftCur])
            leftCur++
        } else if left[leftCur] < right[rightCur] {
            res = append(res, left[leftCur])
            leftCur++
        } else {
            res = append(res, right[rightCur])
            rightCur++
        }
    }
    fmt.Println(res)
    return res
}

