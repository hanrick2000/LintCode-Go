/**
 * @param s: the maximum length of s is 1000
 * @return: the longest palindromic subsequence's length
 */
 
func longestPalindromeSubseq (s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    
    // initilize the dp matrix
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }
    
    for gap := 0; gap < n; gap++ {
        for cur := 0; cur + gap < n; cur++ {
            if gap == 0 {
               dp[cur][cur + gap] = 1 
            } else if gap == 1 && s[cur] == s[cur + gap] {
              dp[cur][cur + gap] =  2
            } else if s[cur] == s[cur + gap] {
                dp[cur][cur + gap] = dp[cur + 1][cur + gap - 1] + 2
            } else if dp[cur + 1][cur + gap] > dp[cur][cur + gap - 1]{
                dp[cur][cur + gap] = dp[cur + 1][cur + gap]
            } else {
                dp[cur][cur + gap] = dp[cur][cur + gap - 1]
            }
        }
    }
    return dp[0][n - 1]
}

