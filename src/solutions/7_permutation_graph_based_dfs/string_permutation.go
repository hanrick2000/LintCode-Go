/**
 * @param A: a string
 * @param B: a string
 * @return: a boolean
 */
 import (
     "strings"
     "sort"
)
 
 
func Permutation (A string, B string) bool {
    stringsA := strings.Split(A, "")
    stringsB := strings.Split(B, "")
    sort.Strings(stringsA)
    sort.Strings(stringsB)
    if len(stringsA) != len(stringsB) {
        return false
    }
    for i := 0; i < len(stringsA); i++ {
        if stringsA[i] != stringsB[i] {
            return false
        }
    }
    return true
}

