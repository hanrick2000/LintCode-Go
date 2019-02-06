func decodeString(s string) string {
    num := 0    
    nums := make([]int, 0) 
    res := make([]byte, 0)
    
    for i := 0; i < len(s); i++ {
        char := s[i]
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            res = append(res, char)
        } else if char == '[' {
            nums = append(nums, num)
            res = append(res, char)
            num = 0
        } else if (char >= '0' && char <= '9') {
            num = num * 10 + int(char - '0')
        } else {
            temp := make([]byte, 0)
            for i := len(res) - 1; i >= 0; i-- {
                if res[i] == '[' {
                    temp = append(temp, res[i + 1:]...)
                    res = res[:i]
                    break
                }
            }
            fmt.Println(temp)
            repeat := nums[len(nums) - 1]
            nums = nums[:len(nums) - 1]
            for i := 0; i < repeat; i++ {
                res = append(res, temp...)
            }
        }
    }
    return string(res)
}
