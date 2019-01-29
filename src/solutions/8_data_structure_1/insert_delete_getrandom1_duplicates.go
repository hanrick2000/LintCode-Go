type RandomizedCollection struct {
    indexMap map[int]map[int]struct{}
    array []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
    return RandomizedCollection {
        make(map[int]map[int]struct{}),
        make([]int, 0),
    }
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
    if _, ok := this.indexMap[val]; !ok {
        this.indexMap[val] = make(map[int]struct{})
    }    
    this.indexMap[val][len(this.array)] = struct{}{}
    this.array = append(this.array, val)
    return len(this.indexMap[val]) == 1
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
    if _, ok := this.indexMap[val]; ok {
        rmIndex := -1
        for k, _ := range this.indexMap[val] {
            rmIndex = k
            break
        }
        delete(this.indexMap[val], rmIndex)
        if len(this.indexMap[val]) == 0 {
            delete(this.indexMap, val)
        }
        
        // remove element
        n := len(this.array)
        this.array[n - 1], this.array[rmIndex] =
            this.array[rmIndex], this.array[n - 1]
        this.array = this.array[:n - 1]
        if rmIndex != n - 1 {
            this.indexMap[this.array[rmIndex]][rmIndex] = struct{}{}
            delete(this.indexMap[this.array[rmIndex]], n - 1)
        }
        return true
    } else {
        return false
    }
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
    return this.array[rand.Intn(len(this.array))]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
