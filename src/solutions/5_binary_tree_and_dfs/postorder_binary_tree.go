/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]TreeNode, 0)
    cur := root
    for len(stack) != 0 || cur != nil {
        if cur != nil {
            res = append(res, cur.Val)
            stack = append(stack, *cur)
            cur = cur.Right
        } else {
            cur = &stack[len(stack) - 1]
            stack = stack[0 : len(stack) - 1]
            cur = cur.Left
        }
    }
    for i := 0; i < len(res) / 2; i++ {
        res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
    }
    return res
}

