import (
    "container/heap"
    "strings"
)

type wordCount struct {
    word string
    count int
}

type wcHeap []wordCount

func (wch wcHeap) Len() int {
    return len(wch)
}

func (wch wcHeap) Less(i int, j int) bool {
    if wch[i].count == wch[j].count {
        return strings.Compare(wch[i].word, wch[j].word) > 0
    }
    return wch[i].count < wch[j].count
}

func (wch wcHeap) Swap(i int, j int) {
    wch[i], wch[j] = wch[j], wch[i]
}

func (wch *wcHeap) Push(x interface{}) {
    *wch = append(*wch, x.(wordCount))
}

func (wch *wcHeap) Pop() interface{} {
    n := len(*wch)
    res := (*wch)[n - 1]
    *wch = (*wch)[:n - 1]
    return res
}

func topKFrequent(words []string, k int) []string {
    wc := make(map[string]int)
    for _, word := range words {
        if _, ok := wc[word]; ok {
            wc[word]++
        } else {
            wc[word] = 1
        }
    }
    
    hp := &wcHeap{}
    for word, count := range wc {
        heap.Push(hp, wordCount{word, count})
        if hp.Len() > k {
            heap.Pop(hp)
        }
    }
    
    res := make([]string, k)
    for i := k - 1; i >= 0; i-- {
        res[i] = heap.Pop(hp).(wordCount).word
    }
    return res
}
