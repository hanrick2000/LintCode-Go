/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    res := root.Val
    cur := root 
    for cur != nil {
        res = closer(res, cur.Val, target)
        if float64(cur.Val) < target {
            cur = cur.Right
        } else {
            cur = cur.Left
        }
    }
    return res
}

func closer(val1 int, val2 int, target float64) int {
    if abs(float64(val1) - target) < abs(float64(val2) - target) {
        return val1
    } else {
        return val2
    }
}

func abs(num float64) float64 {
    if num < 0 {
        return -num
    } else {
        return num
    }
}
