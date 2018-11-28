/**
 * @param A: an integer array
 * @param target: An integer
 * @param k: An integer
 * @return: an integer array
 */
func kClosestNumbers (A []int, target int, k int) []int {
    // write your code here
    res := make([]int, 0)
    if k == 0 {
        return res
    }
    closestIndex := getClosestIndex(A, target)
    res = append(res, A[closestIndex])
    left, right := closestIndex - 1, closestIndex + 1
    for len(res) < k {
        if left < 0 {
            res = append(res, A[right])
            right++
        } else if right >= len(A) {
            res = append(res, A[left])
            left--
        } else if abs(A[left] - target) <= abs(A[right] - target) {
            res = append(res, A[left])
            left--
        } else {
            res = append(res, A[right])
            right++
        }
    }
    return res
}

func getClosestIndex(A []int, target int) int {
    left, right := 0, len(A) - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if A[mid] < target {
            left = mid
        } else {
            right = mid
        }
    }
    if abs(A[left] - target) <= abs(A[right] - target) {
        return left
    } else {
        return right
    }
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
