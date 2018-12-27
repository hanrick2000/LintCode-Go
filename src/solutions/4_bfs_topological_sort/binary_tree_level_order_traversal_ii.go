/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
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
        res = append(res, level)
    }
    
    for i := 0; i < len(res) / 2; i++ {
        res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
    }
    return res
}
