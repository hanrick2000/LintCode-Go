/**
 * @param S: A list of integers
 * @return: An integer
 */
import "sort"

func triangleCount (S []int) int {
    res := 0
    n := len(S)
    sort.Ints(S)    
    
    for mid := 1; mid < n - 1; mid++ {
        left, right := 0, mid + 1
        for left < mid && right < n {
            if S[right] < S[mid] + S[left] {
                right++
                res += mid - left
            } else {
                left++
            }
        }
    }
    return res
}

// 3, 4, 6, 7
// left=0, mid=1, right=1 res = 0
// left=0, mid=1, right=2 
// left=0, mid=1, right=3 res = 1
// left=0, mid=2, right=4 res = 2
// left=1, mid=2, right=3
// left=1, mid=2, right=4 res = 3
// left=2, mid=3, right=3
// left=2, mid=3, right=4
// res=3

