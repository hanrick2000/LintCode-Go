func helper(A []int, target int, leftmost bool) int {
    if len(A) == 0 {
        return -1
    }
    left, right := 0, len(A) - 1
    for left + 1 < right {
        mid := left + (right - left) / 2
        if A[mid] < target {
            left = mid
        } else if A[mid] > target {
            right = mid
        } else if leftmost {
            right =  mid
        } else {
            left = mid
        }
    }
    if leftmost {
        if A[left] == target {
            return left
        } else if A[right] == target {
            return right
        } else {
            return -1
        }
    } else {
        if A[right] == target {
            return right
        } else if A[left] == target {
            return left
        } else {
            return -1
        }
    }
}
 
func searchRange (A []int, target int) []int {
    return []int{helper(A, target, true), helper(A, target, false)}
}

