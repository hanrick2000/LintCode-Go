/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    fast := head
    slow := head
    
    for fast != nil && slow != nil {
        slow = slow.Next
        if fast.Next == nil {
            return false
        }
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    } 
    
    return false
}
