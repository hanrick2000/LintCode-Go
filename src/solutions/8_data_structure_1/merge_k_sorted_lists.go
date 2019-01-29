/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"

type listHeap []*ListNode

func (lh listHeap) Len() int {
    return len(lh)
}

func (lh listHeap) Less(i int, j int) bool {
    return lh[i].Val < lh[j].Val
}

func (lh listHeap) Swap(i int, j int) {
    lh[i], lh[j] = lh[j], lh[i]
}

func (lh *listHeap) Pop() interface{} {
    n := len(*lh)
    res := (*lh)[n - 1]
    *lh = (*lh)[:n - 1]
    return res
}

func (lh *listHeap) Push(node interface{}) {
    *lh = append(*lh, node.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
    res := &ListNode{0, nil}
    lh := &listHeap{}
    for _, list := range lists {
        if list != nil {
            heap.Push(lh, list)
        }
    }
    
    cur := res
    for lh.Len() > 0 {
        curPop := heap.Pop(lh).(*ListNode)
        cur.Next = curPop
        if curPop.Next != nil {
            heap.Push(lh, curPop.Next)
        }
        cur = cur.Next
    }
    
    return res.Next
}
