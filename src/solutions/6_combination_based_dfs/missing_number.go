func missingNumber(nums []int) int {
    n := len(nums)
    res := n * (n + 1) / 2
    for _, num := range nums {
        res -= num
    }
    return res
}
