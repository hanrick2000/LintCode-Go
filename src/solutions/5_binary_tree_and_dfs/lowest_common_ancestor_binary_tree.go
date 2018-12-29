/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    _, res := helper(root, p, q)
    return res
}

func helper(root, p, q *TreeNode) (bool, *TreeNode) {
    if root == nil {
        return false, nil
    }
    leftContains, leftRes := helper(root.Left, p, q)
    rightContains, rightRes := helper(root.Right, p, q)
    fmt.Println(root.Val)
    if leftRes != nil {
        return true, leftRes
    }
    if rightRes != nil {
        return true, rightRes 
    }
    if root.Val == p.Val || root.Val == q.Val {
        if leftContains || rightContains {
            return true, root
        } else {
            return true, nil
        }
    }
    if leftContains && rightContains {
        return true, root
    }
    if leftContains || rightContains {
        return true, nil
    }

    return false, nil
}
