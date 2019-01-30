type Vector2D struct {
    nums [][]int
    row int
    col int
}


func Constructor(v [][]int) Vector2D {
    res := Vector2D{v, 0, 0}
    res.moveToNext()
    return res
}


func (this *Vector2D) Next() int {
    res := this.nums[this.row][this.col]
    this.col++
    this.moveToNext()
    return res
}

func (this *Vector2D) moveToNext() {
    for this.row < len(this.nums) && len(this.nums[this.row]) == this.col {
        this.row++
        this.col = 0
    }
}


func (this *Vector2D) HasNext() bool {
    return this.row < len(this.nums)
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
