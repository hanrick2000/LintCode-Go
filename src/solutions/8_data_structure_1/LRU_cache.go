// Creating doubly linked list
type ListNode struct {
    key int
    val int
    prev *ListNode
    next *ListNode
}

type list struct {
    head *ListNode
    tail *ListNode
}

func (l *list) removeNode(node *ListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (l *list) addNode(node *ListNode) {
    node.next = l.head.next
    node.prev = l.head
    node.next.prev = node
    l.head.next = node
}


func NewList() list {
    l := list{&ListNode{0, 0, nil, nil}, &ListNode{0, 0, nil, nil}}
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

// LRU cache
type LRUCache struct {
    l list
    values map[int]*ListNode
    capacity int
}


func Constructor(capacity int) LRUCache {
    l := NewList()
    return LRUCache{l, make(map[int]*ListNode), capacity}
}


func (this *LRUCache) Get(key int) int {
    if targetNode, ok := this.values[key]; ok {
        this.l.removeNode(targetNode)
        this.l.addNode(targetNode)
        return targetNode.val
    } else {
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if targetNode, ok := this.values[key]; ok {
        targetNode.val = value
        this.l.removeNode(targetNode)
        this.l.addNode(targetNode)
    } else {
        add := &ListNode{key, value, nil, nil}
        this.values[key] = add
        this.l.addNode(add)
        if len(this.values) > this.capacity {
            removeKey := this.l.tail.prev.key
            this.l.removeNode(this.l.tail.prev)
            delete(this.values, removeKey)
        }
    }
}

