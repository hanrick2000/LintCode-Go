/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]TreeNode, 0)
    cur := root
    for cur != nil || len(stack) != 0 {
        if cur != nil {
            res = append(res, cur.Val)
            stack = append(stack, *cur)
            cur = cur.Left
        } else {
            cur = &stack[len(stack) - 1]
            stack = stack[0 : len(stack) - 1]
            cur = cur.Right
        }
    }
    return res
}
