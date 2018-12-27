/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    opposite := false
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    queue := make([]TreeNode, 0)
    queue = append(queue, *root)
    for len(queue) > 0 {
        size := len(queue)
        level := make([]int, 0)
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            level = append(level, cur.Val)
            if cur.Left != nil {
                queue = append(queue, *cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, *cur.Right)
            }
        }
        if opposite {
            for i := 0; i < len(level) / 2; i++ {
                level[i], level[len(level) - i - 1] = level[len(level) - i - 1], level[i]
            }
        }
        opposite = !opposite
        res = append(res, level)
    }
    return res
}
