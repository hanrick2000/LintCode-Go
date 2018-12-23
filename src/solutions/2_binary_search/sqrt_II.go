/**
 * @param x: a double
 * @return: the square root of x
 */
func sqrt (x float64) float64 {
    // write your code here
    var left, right float64 = 0, x
    if right < 1 {
        right = 1
    }
    for right - left > 0.00001 {
        mid := left + (right - left) / 2
        if mid < x / mid {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}

