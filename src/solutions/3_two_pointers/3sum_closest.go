/**
 * @param numbers: Give an array numbers of n integer
 * @param target: An integer
 * @return: return the sum of the three integers, the sum closest target.
 */
import "sort"

func threeSumClosest (numbers []int, target int) int {
    n := len(numbers)
    sort.Ints(numbers)
    res := numbers[0] + numbers[1] + numbers[2]
    for i := 0; i < n; i++ {
        left, right := i + 1, n - 1
        for left < right {
            cur := numbers[left] + numbers[right] + numbers[i]
            res = closerNum(res, cur, target)
            if cur < target {
                left++
            } else {
                right--
            }
        }
    }
    return res
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func closerNum(num1 int, num2 int, target int) int {
    if abs(num1 - target) < abs(num2 - target) {
        return num1
    }
    return num2
}
