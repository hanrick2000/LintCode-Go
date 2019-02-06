func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }
    
    res := 0
    for _, num := range nums {
        if _, ok := numSet[num]; ok {
            temp := getConsecutive(numSet, num)
            if res < temp {
                res = temp
            }
        }
    }
    return res
}

func getConsecutive(numSet map[int]struct{}, num int) int {
    res := 0
    cur := num
    for true {
        if _, ok := numSet[cur]; ok {
            res++
            delete(numSet, cur)
            cur--
        } else {
            break
        }
    }
    cur = num + 1
    for true {
        if _, ok := numSet[cur]; ok {
            res++
            delete(numSet, cur)
            cur++
        } else {
            break
        }
    }
    return res
}
