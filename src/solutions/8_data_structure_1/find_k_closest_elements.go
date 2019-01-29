import "container/heap"

type maxHeap struct {
    nums []int
    x int
}

func (h maxHeap) Len() int {
    return len(h.nums)
}

func (h maxHeap) Less(i, j int) bool {
    if abs(h.nums[j] - h.x) == abs(h.nums[i] - h.x) {
        return h.nums[j] < h.nums[i]
    }
    return abs(h.nums[j] - h.x) < abs(h.nums[i] - h.x)
}

func (h maxHeap) Swap(i, j int) {
    h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *maxHeap) Pop() interface{} {
    n := len(h.nums)
    num := h.nums[n - 1]
    h.nums = h.nums[: n - 1]
    return num
}

func (h *maxHeap) Push(x interface{}) {
    h.nums = append(h.nums, x.(int))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func findClosestElements(arr []int, k int, x int) []int {
    hp := &maxHeap{make([]int, 0), x}
    for _, num := range arr {
        heap.Push(hp, num)
        if hp.Len() > k {
            heap.Pop(hp)
        }
    }
    
    res := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        res[i] = heap.Pop(hp).(int)
    }
    sort.Ints(res)
    return res
}

