func nextPermutation(nums []int)  {
    if len(nums) == 0 {
        return
    }
    n := len(nums)
    max := nums[n - 1]
    for cur := n - 2; cur >= 0; cur-- {
        if nums[cur] < max {
            // We get a number to change
            replaceIndex := cur + 1
            for i := cur + 1; i < n; i++ {
                if nums[i] > nums[cur] && nums[i] < nums[replaceIndex] {
                   replaceIndex = i
                }
            }
            nums[cur], nums[replaceIndex] = nums[replaceIndex], nums[cur]
            sort.Ints(nums[cur + 1:])
            return
        } else {
            max = nums[cur]
        }
    }
    
    sort.Ints(nums)
    return
}
