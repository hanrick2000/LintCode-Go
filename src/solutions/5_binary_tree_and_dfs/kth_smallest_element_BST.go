/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var cur int
var res int

func kthSmallest(root *TreeNode, k int) int {
    cur = 0
    dfs(root, k)
    return res
}

func dfs(root *TreeNode, k int) {
    if root == nil {
        return
    }
    dfs(root.Left, k)
    cur++
    if cur == k {
        res = root.Val
    }
    dfs(root.Right, k)
}
