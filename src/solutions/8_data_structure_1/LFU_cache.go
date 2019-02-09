type listNode struct {
    key int
    num int
    freq int
    prev *listNode
    next *listNode
}

type linkedlist struct {
    head *listNode  
    tail *listNode
}

func newLinkedList() *linkedlist {
    head := &listNode{0, 0, 0, nil, nil}
    tail := &listNode{0, 0, 0, nil, nil}
    head.next = tail
    tail.prev = head
    return &linkedlist{head, tail}
}

func (l *linkedlist) empty() bool {
    return l.head.next == l.tail
}

func (l *linkedlist) add(node *listNode) {
    node.prev = l.head
    node.next = l.head.next
    node.next.prev = node
    node.prev.next = node
}

func (l *linkedlist) remove(node *listNode) {
    node.next.prev = node.prev
    node.prev.next = node.next
}

type LFUCache struct {
    values map[int]*listNode
    freqList map[int]*linkedlist
    minFreq int
    capacity int
}

func Constructor(capacity int) LFUCache {
    return LFUCache {
        make(map[int]*listNode),
        make(map[int]*linkedlist),
        0, capacity,
    }
}

func(this *LFUCache) remove(node *listNode) {
    freq := node.freq
    this.freqList[freq].remove(node)
    if this.freqList[freq].empty() {
        delete(this.freqList, freq)
    }
    delete(this.values, node.key)
}

func(this *LFUCache) add(node *listNode) {
    freq := node.freq
    if _, ok := this.freqList[freq]; !ok {
        this.freqList[freq] = newLinkedList()
    }
    this.freqList[freq].add(node)
    this.values[node.key] = node
}

func (this *LFUCache) Get(key int) int {
    if res, ok := this.values[key]; ok {
        // remove res from current freq linkedlist
        this.remove(res)
        
        // Add to the new freq linkedList
        res.freq++
        this.add(res)
        
        return res.num
    }
    return -1
}

func (this *LFUCache) Put(key int, value int)  {
    if this.capacity == 0 {
        return
    }
    if node, ok := this.values[key]; ok {
        node.num = value
        this.Get(key)
    } else {
        // Evict element if capacity reaches limit
        if len(this.values) == this.capacity {
            for  {
                if _, ok := this.freqList[this.minFreq]; ok {
                    break
                }
                this.minFreq++
            }
            tempList := this.freqList[this.minFreq]
            delete(this.values, tempList.tail.prev.key)
            tempList.remove(tempList.tail.prev)
        }
        
        // Add new node if not exists
        this.minFreq = 1
        newNode := &listNode{key, value, 1, nil, nil}
        this.add(newNode)
    }
}

