/**
 * @param nums: a continuous stream of numbers
 * @param number: a number
 * @return: returns the first unique number
 */
func firstUniqueNumber (nums []int, number int) int {
    // Write your code here
    count := make(map[int]int)
    found := false
    for _, num := range nums {
        if _, ok := count[num]; ok {
            count[num]++
        } else {
            count[num] = 1
        }
        if num == number {
            found = true
            break
        }
    }
    if !found {
        return -1
    }
    for _, num := range nums {
        if count[num] == 1 {
            return num
        }
    }
    return -1
    
}

