/**
 * @param n: An integer
 * @param str: a string with number from 1-n in random order and miss one number
 * @return: An integer
 */
func findMissing2 (n int, str string) int {
    res := make([]int, 0)
    candidate := make([]int, 0)
    helper(n, &candidate, &res, str, 0)
    sum := n * (n + 1) / 2
    for _, num := range res {
        sum -= num
    }
    return sum
}

func helper(n int, candidate *[]int, res *[]int, 
                str string, start int) {
    if start == len(str) {
        if len(*candidate) == n - 1 {
            *res = make([]int, n - 1)
            copy(*res, *candidate)
        }
        return
    }
    
    if len(*res) == n - 1 {
        return
    }
    
    if str[start] == '0' {
        return
    }
    
   if !contains(*candidate, int(str[start] - '0')) {
       *candidate = append(*candidate, int(str[start] - '0'))
       helper(n, candidate, res, str, start + 1)
       *candidate = (*candidate)[0 : len(*candidate) - 1]
   }
   
   if start + 1 < len(str) {
       num := int(str[start] - '0') * 10 + int(str[start + 1] - '0')
       if num <= n && !contains(*candidate, num) {
           *candidate = append(*candidate, num)
            helper(n, candidate, res, str, start + 2)
            *candidate = (*candidate)[0 : len(*candidate) - 1]
       }
   }
}

func contains(nums []int, target int) bool {
    for _, num := range nums {
        if num == target {
            return true
        }
    }
    return false
}
