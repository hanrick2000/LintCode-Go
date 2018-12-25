/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for slow != nil && fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    
    if slow == nil || fast == nil {
        return nil
    }
    
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
