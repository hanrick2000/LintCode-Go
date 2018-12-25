func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    curA := headA
    curB := headB
    firstPassA := true
    firstPassB := true
    
    for curA != curB {
        curA = curA.Next
        curB = curB.Next
        if curA == nil {
            if firstPassA {
                firstPassA = !firstPassA
                curA = headB
            } else {
                return nil
            }
        }
        if curB == nil {
            if firstPassB {
                firstPassB = !firstPassB
                curB = headA
            } else {
                return nil
            }
        }
    }
    return curA
}
