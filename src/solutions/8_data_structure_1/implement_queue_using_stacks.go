type MyQueue struct {
    sk1 []int
    sk2 []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{make([]int, 0), make([]int, 0)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.sk2 = append(this.sk2, x)
}


func (this *MyQueue) DumpSk1ToSk2() {
    for len(this.sk2) > 0 {
        num := this.sk2[len(this.sk2) - 1]
        this.sk1 = append(this.sk1, num)
        this.sk2 = this.sk2[:len(this.sk2) - 1]
    }
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.sk1) == 0 {
        this.DumpSk1ToSk2()
    }
    res := this.sk1[len(this.sk1) - 1]
    this.sk1 = this.sk1[:len(this.sk1) - 1] 
    return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.sk1) == 0 {
        this.DumpSk1ToSk2()
    }
    return this.sk1[len(this.sk1) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.sk1) == 0 && len(this.sk2) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
