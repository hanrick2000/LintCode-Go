/**
 * @param str: A string
 * @return: all permutations
 */
 import "sort"
 
 
func stringPermutation2 (str string) []string {
    strByte := []byte(str)
    res := make([]string, 0)
    helper(&strByte, 0, &res)
    sort.Strings(res)
    return res
}

func helper(str *[]byte, cur int, res *[]string) {
    if cur == len(*str) {
        *res = append(*res, string(*str))
        return
    }
    
    visited := make(map[byte]struct{})
    for i := cur; i < len(*str); i++ {
        if _, ok := visited[(*str)[i]]; !ok {
            (*str)[i], (*str)[cur] = (*str)[cur], (*str)[i]
            visited[(*str)[cur]] = struct{}{}
            helper(str, cur + 1, res)
            (*str)[cur], (*str)[i] = (*str)[i], (*str)[cur]
        }
    }
}

