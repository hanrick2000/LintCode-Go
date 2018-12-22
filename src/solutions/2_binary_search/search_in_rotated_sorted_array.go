func search (A []int, target int) int {
    // write your code here
    n := len(A)
    if n == 0 {
        return -1
    }
    left, right := 0, n - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if A[mid] > A[0] {
            if target >= A[left] && target <= A[mid] {
                right = mid
            } else {
                left = mid
            }
        } else {
            if target >= A[mid] && target <= A[right] {
                left = mid
            } else {
                right = mid
            }
        }
    }
    if A[left] == target {
        return left
    } else if A[right] == target {
        return right
    } else {
        return -1
    }
}
