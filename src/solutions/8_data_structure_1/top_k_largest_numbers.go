/**
 * @param nums: an integer array
 * @param k: An integer
 * @return: the top k largest numbers in array
 */
import "container/heap"
 
type minHeap []int

func (h minHeap) Len() int {
    return len(h)
}

func (h minHeap) Less(i int, j int) bool {
    return h[i] < h[j]   
}

func (h minHeap) Swap(i int, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(num interface{}) {
    *h = append(*h, num.(int))
}

func (h *minHeap) Pop() interface{} {
    n := len(*h)
    num := (*h)[n - 1]
    *h = (*h)[:n - 1]
    return num
}
 
func topk (nums []int, k int) []int {
    if k <= 0 {
        return make([]int, 0)
    }
    if k > len(nums) {
        k = len(nums)
    }
    
    mhp := &minHeap{}  
    for _, num := range nums {
        heap.Push(mhp, num)
        if mhp.Len() > k {
            heap.Pop(mhp)
        }
    }
    
    res := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        res[i] = heap.Pop(mhp).(int)
    }
    return res
}

