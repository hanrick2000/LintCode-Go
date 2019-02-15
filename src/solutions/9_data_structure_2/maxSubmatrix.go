func maxSubmatrix (matrix [][]int) int {
    // write your code here
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    res := make([][]int, 0)
    for i := 0; i <= len(matrix); i++ {
        res = append(res, make([]int, len(matrix[0]) + 1))
    }
    for i := 1; i < len(res); i++ {
        for j := 1; j < len(res[0]); j++ {
            res[i][j] = res[i][j - 1] + res[i - 1][j] + matrix[i - 1][j - 1] - res[i - 1][j - 1]
        }
    }
    
    maxNum := 0
    maxIndex := []int{-1, -1, -1, -1}
    for i := 0; i < len(res); i++ {
        for j := 0; j < len(res[0]); j++ {
            for m := i + 1; m < len(res); m++ {
                for n := j + 1; n < len(res[0]); n++ {
                   curSum := res[m][n] - res[m][j] - res[i][n] + res[i][j]
                   if maxIndex[0] == -1 || curSum > maxNum {
                        maxNum = curSum
                        maxIndex = []int{0, 0, 0, 0}
                   }
                }
            }
        }
    }
    
   return maxNum 
}
