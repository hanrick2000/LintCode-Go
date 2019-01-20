func nextClosestTime(time string) string {
    nums := make([]int, 0)
    
    for _, temp := range time {
        if temp == ':' {
            continue
        }
        nums = append(nums, int(temp - '0'))
    }
    res := make([]int, len(nums))
    copy(res, nums)

     // Nothing found
    min := nums[0]
    for _, num := range nums {
        if num < min {
            min = num
        }
    }
    
    for i := 3; i >= 0; i-- {
        replace := nextLargerNum(res[i], nums)
        if replace != -1 {
            res[i] = replace
            if isValid(res) {
                for j := i + 1; j < 4; j++ {
                    res[j] = min
                }
                return parseResult(res)
            }
            res[i] = nums[i]
        }
    }
    
    return parseResult([]int{min, min, min, min})
}

func nextLargerNum(target int, nums []int) int {
    res := -1
    for _, num := range nums {
        if num > target {
            if res == -1 || res > num {
                res = num
            }
        }
    }
    return res
}

func isValid(nums []int) bool {
    if nums[0] * 10 + nums[1] > 23 {
        return false
    }
    if nums[2] * 10 + nums[3] > 59 {
        return false
    }
    return true
}

func parseResult(nums []int) string {
    num1 := strconv.FormatInt(int64(nums[0] * 10 + nums[1]), 10)
    if len(num1) < 2 {
        num1 = "0" + num1
    }
    num2 := strconv.FormatInt(int64(nums[2] * 10 + nums[3]), 10)
    if len(num2) < 2 {
        num2 = "0" + num2
    }
    return num1 + ":" + num2
}
