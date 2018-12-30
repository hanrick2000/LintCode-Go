/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func closestKValues(root *TreeNode, target float64, k int) []int {
    res := make([]int, 0)
    helper(root, target, k, &res)
    return res
}

func helper(root *TreeNode, target float64, k int, res *[]int) {
    if root == nil {
        return
    }
    helper(root.Left, target, k, res)
    if len(*res) < k {
        *res = append(*res, root.Val)
    } else if float64(root.Val) <= target {
        *res = (*res)[1:]
        *res = append(*res, root.Val)
    } else if closer(target, root.Val, (*res)[0]) {
        *res = (*res)[1:]
        *res = append(*res, root.Val)
    }
    helper(root.Right, target, k, res)
}

func closer(target float64, num1 int, num2 int) bool {
    if abs(target - float64(num1)) < abs(target - float64(num2)) {
        return true
    }
    return false
}

func abs(num float64) float64 {
    if num < 0 {
        return -num
    }
    return num
}
