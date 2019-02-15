type NumArray struct {
   sum []int 
}


func Constructor(nums []int) NumArray {
    sum := make([]int, len(nums) + 1)
    res := NumArray{sum}
    for i, num := range nums {
        res.Update(i, num)
    }
    return res
}


func (this *NumArray) Update(i int, val int)  {
    cur := i + 1
    curVal := this.sumStart(i) - this.sumStart(i - 1)
    diff := val - curVal
    
    for cur < len(this.sum) {
        this.sum[cur] += diff
        cur += (cur & -cur)
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.sumStart(j) - this.sumStart(i - 1)
}

func (this *NumArray) sumStart(i int) int {
    res := 0
    i++
    for i > 0 {
        res += this.sum[i]
        i -= (i & -i)
    }
    return res
}


