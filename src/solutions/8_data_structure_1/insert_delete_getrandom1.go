type RandomizedSet struct {
    indexMap map[int]int
    array []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        make(map[int]int),
        make([]int, 0),
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.indexMap[val]; !ok {
        this.indexMap[val] = len(this.array)
        this.array = append(this.array, val)
        return true
    } else {
        return false
    }
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.indexMap[val]; ok {
        rmIndex := this.indexMap[val]
        n := len(this.array)
        this.array[rmIndex], this.array[n - 1] = 
            this.array[n - 1], this.array[rmIndex]
        this.array = this.array[0 : n - 1]
        delete(this.indexMap, val)
        if rmIndex != n - 1 {
            this.indexMap[this.array[rmIndex]] = rmIndex
        }
        return true
    } else {
        return false
    }
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    return this.array[rand.Intn(len(this.array))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
