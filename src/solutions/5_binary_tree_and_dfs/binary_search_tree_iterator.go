/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    res := BSTIterator{make([]TreeNode, 0)}
    cur := root
    for cur != nil {
        res.stack = append(res.stack, *cur)
        cur = cur.Left
    }
    return res
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    last := this.stack[len(this.stack) - 1]
    this.stack = this.stack[0 : len(this.stack) - 1]
    cur := last.Right
    for cur != nil {
        this.stack = append(this.stack, *cur)
        cur = cur.Left
    }
    return last.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
