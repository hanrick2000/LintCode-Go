/**
 * @param colors: A list of integer
 * @param k: An integer
 * @return: nothing
 */
func sortColors2 (colors *[]int, k int)  {
    // write your code here
    for i := 1; i <= k; i++ {
        helper(colors, i)
    }
}

func helper(colors *[]int, pivot int) {
    prev := 0
    for i, val := range *colors {
        if val <= pivot {
            (*colors)[i], (*colors)[prev] = (*colors)[prev], (*colors)[i]
            prev++
        }
    }
}
