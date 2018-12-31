/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: the root of binary tree.
 * @return: An integer
 */

var res int = -2147483648 
 
func maxPathSum2 (root *TreeNode) int {
    // write your code here
    res = -2147483648
    helper(root, 0)
    return res
}

func helper(root *TreeNode, curSum int) {
    if root == nil {
        return
    }
    curSum += root.Val
    if curSum > res {
        res = curSum
    }
    helper(root.Left, curSum)
    helper(root.Right, curSum)
}

