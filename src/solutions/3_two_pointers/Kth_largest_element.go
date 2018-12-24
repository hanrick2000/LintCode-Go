/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 */
 import "container/heap"
 
func kthLargestElement (n int, nums []int) int {
    // Solution 1: with heap
    //return solutionHeap(n, nums)
    
    // Solution 2: with recursion
    return quickSearch(len(nums) - n + 1, nums)
}

func quickSearch(n int, nums []int) int {
    pivot := nums[len(nums) - 1]
    prev := 0
    
    for i, val := range nums {
        if val <= pivot {
            nums[i], nums[prev] = nums[prev], nums[i]
            prev++
        }
    }
    prev--
    if prev + 1 == n {
        return pivot
    } else if prev + 1 > n {
        return quickSearch(n, nums[0 : prev])
    } else {
        return quickSearch(n - prev - 1, nums[prev + 1:])
    }
}

type minHeap []int

func (h *minHeap) Len() int {
    return len(*h)
}

func (h *minHeap) Less(i int, j int) bool {
    return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i int, j int) {
    (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(val interface{}) {
    *h = append(*h, val.(int))
}

func (h *minHeap) Pop() interface{} {
    n := len(*h)
    val := (*h)[n - 1]
    *h = (*h)[0 : n - 1]
    return val
}

func solutionHeap(n int, nums []int) int {
    h := &minHeap{}
    heap.Init(h)
    for _, num := range nums {
        heap.Push(h, num)
        if len(*h) > n {
            heap.Pop(h)
        }
    }
    return heap.Pop(h).(int)
}

