func nthUglyNumber(n int) int {
    nums := make([][]int, 3)
    for i := 0; i < 3; i++ {
        nums[i] = make([]int, 0)
    }
    nums[0] = append(nums[0], 1)
    popNum := 0
    i := 0
    for i < n {
        popIndex := minIndex(&nums)
        nextNum := nums[popIndex][0]
        nums[popIndex] = nums[popIndex][1:]
        
        if nextNum != popNum {
            i++
            popNum = nextNum
            nums[0] = append(nums[0], 2 * nextNum)
            nums[1] = append(nums[1], 3 * nextNum)
            nums[2] = append(nums[2], 5 * nextNum)
        }
        
    }
    return popNum
}

func minIndex(nums *[][]int) int {
    min := (*nums)[0][0]
    ind := 0
    
    for i := 0; i < 3; i++ {
        if len((*nums)[i]) == 0 {
            continue
        }
        if (*nums)[i][0] < min {
            min = (*nums)[i][0]
            ind = i
        }
    }
    return ind
}
