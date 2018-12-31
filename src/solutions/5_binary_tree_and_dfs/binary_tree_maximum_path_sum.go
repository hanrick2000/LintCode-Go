/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: The root of binary tree.
 * @return: An integer
 */
func maxPathSum (root *TreeNode) int {
    // write your code here
    _, res := helper(root)
    return res
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, -2147483648 
    }
    leftLocal, leftGlobal := helper(root.Left)
    rightLocal, rightGlobal := helper(root.Right)
    local := root.Val
    local += max([]int{0, leftLocal, rightLocal})
    global := root.Val
    if leftLocal > 0 {
        global += leftLocal
    }
    if rightLocal > 0 {
        global += rightLocal
    }
    global = max([]int{global, leftGlobal, rightGlobal})
    return local, global
}

func max(nums []int) int {
    res := nums[0]
    for _, num := range nums {
        if num > res {
            res = num
        }
    }
    return res
}
