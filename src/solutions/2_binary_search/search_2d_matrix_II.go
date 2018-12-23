func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    row, col := len(matrix) - 1, 0
    for row >= 0 && col <= len(matrix[0]) - 1 {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            col++
        } else {
            row--
        }
    }
    return false
}
