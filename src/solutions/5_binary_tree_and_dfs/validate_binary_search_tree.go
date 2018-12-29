/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    _, _, res := helper(root)
    return res
}

func helper(root *TreeNode) (int, int, bool) {
    max, min := root.Val, root.Val
    if root.Left != nil {
        curMax, curMin, isbst := helper(root.Left)
        if !isbst {
            return 0, 0, false
        }
        if curMax >= root.Val {
            return 0, 0, false
        } else {
            min = curMin
        }
    }
    if root.Right != nil {
        curMax, curMin, isbst := helper(root.Right)
        if !isbst {
            return 0, 0, false
        }
        if curMin <= root.Val {
            return 0, 0, false
        } else {
            max = curMax
        }
    }
    return max, min, true
}
