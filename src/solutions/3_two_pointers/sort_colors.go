/**
 * @param nums: A list of integer which is 0, 1 or 2 
 * @return: nothing
 */
func helper (nums *[]int, pivot int) {
    prev := 0
    for i, value := range *nums {
        if value <= pivot {
            (*nums)[prev], (*nums)[i] = (*nums)[i], (*nums)[prev]
            prev++
        }
    }
    return
}
 
func sortColors (nums *[]int)  {
    // write your code here
    helper(nums, 0)
    helper(nums, 1)
    return
}

