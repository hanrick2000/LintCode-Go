/**
 * @param pages: an array of integers
 * @param k: An integer
 * @return: an integer
 */
 
func canFinish(pages []int, k int, target int) bool {
    cur := 0
    num := 0
    
    for _, val := range pages {
        if val > target {
            return false
        }
        if cur + val > target {
            num++
            cur = val
        } else {
            cur += val
        }
    }
    num++
    return num <= k
} 
 
func copyBooks (pages []int, k int) int {
    if len(pages) == 0 {
        return 0
    }
    sum := 0
    for _, val := range pages {
        sum += val
    }
    left, right := pages[0], sum
    for left + 1 < right {
        mid := left + (right - left) / 2
        if canFinish(pages, k, mid) {
            right = mid
        } else {
            left = mid
        }
    }
    if canFinish(pages, k, left) {
        return left
    } else {
        return right
    }
}

