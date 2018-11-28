func myPow(x float64, n int) float64 {
    var power int64 = int64(n)
    if power < 0 {
        power = -power
    }
    var res float64 = 1
    var curPower int64 = 1
    helpNumber := x
    
    for i := 0; i < 32; i++ {
        if (curPower & power) > 0 {
            res *= helpNumber
        } 
        curPower = curPower << 1
        helpNumber = helpNumber * helpNumber
    }
    
    if n < 0 {
        return 1.0/res
    } else {
        return res
    }
}
