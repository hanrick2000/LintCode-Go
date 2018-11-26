import "sort"

func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    if nums == nil || len(nums) == 0 {
        return res
    }
    n := len(nums)
    sort.Ints(nums)
    helper(&res, nums, 0, n)
    return res
}

func helper(res *[][]int, nums []int, cur int, n int) {
    if cur == n {
        temp := make([]int, n)
        copy(temp, nums)
        *res = append(*res, temp)
        return
    }
    visited := make(map[int]struct{})
    for i := cur; i < n; i++ {
        if _, exists := visited[nums[i]]; exists {
            continue
        }
        visited[nums[i]] = struct{}{}
        nums[cur], nums[i] = nums[i], nums[cur]
        helper(res, nums, cur+1, n)
        nums[i], nums[cur] = nums[cur], nums[i]
    }
}

