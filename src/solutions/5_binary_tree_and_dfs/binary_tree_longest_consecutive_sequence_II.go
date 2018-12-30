/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
    _, _, res := helper(root)
    return res
}

// return local increase, local decrease and global
func helper(root *TreeNode) (int, int, int) {
    if root == nil {
        return 0, 0, 0
    }
    leftInc, leftDec, leftGlobalMax := helper(root.Left)
    rightInc, rightDec, rightGlobalMax := helper(root.Right)
    
    // Get localInc and localDec
    localInc, localDec := 1, 1
    if root.Left != nil && root.Val == root.Left.Val + 1 {
        localInc = 1 + leftInc
    }
    if root.Right != nil && root.Val == root.Right.Val + 1 {
        localInc = max([]int{1 + rightInc, localInc}) 
    }
    if root.Left != nil && root.Val == root.Left.Val - 1 {
        localDec = 1 + leftDec
    }
    if root.Right != nil && root.Val == root.Right.Val - 1 {
        localDec = max([]int{1 + rightDec, localDec})
    }
    
    // Get globalMax
    globalMax := max([]int{localInc, localDec})
    if root.Left != nil && root.Right != nil {
        if root.Val == root.Left.Val + 1 && root.Val == root.Right.Val - 1 {
            globalMax = 1 + leftInc + rightDec
        } else if root.Val == root.Left.Val - 1 && root.Val == root.Right.Val + 1 {
            globalMax = 1 + leftDec + rightInc
        }
    }
    return localInc, localDec, max([]int{globalMax, leftGlobalMax, rightGlobalMax})
}

func max(nums []int) int {
    res := nums[0]
    for _, num := range nums {
        if res < num {
            res = num
        }
    }
    return res
}

