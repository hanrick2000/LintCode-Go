/**
 * @param dividend: the dividend
 * @param divisor: the divisor
 * @return: the result
 */
func divide (dividend int, divisor int) int {
    // write your code here
    var longDividend int64 = int64(dividend)
    var longDivisor int64 = int64(divisor)
    var isPositive bool = true
    const MAX_NUM = 2147483647

    if divisor == 0 || (divisor == -1 && dividend == MAX_NUM) || (divisor == -1 && dividend == -2147483648){
        return MAX_NUM
    }
    
    if longDividend < 0 {
        longDividend = -longDividend
        isPositive = !isPositive
    }
    if longDivisor <= 0 {
        longDivisor = -longDivisor
        isPositive = !isPositive
    }
    
    res := 0
    curRes := 1
    curNum := longDivisor
    for curNum < longDividend {
        curRes = curRes << 1
        curNum = curNum << 1
    }
    for curRes > 0 {
        if longDividend >= curNum {
            longDividend -= curNum
            res += curRes
        }
        curRes = curRes >> 1
        curNum = curNum >> 1
    }
    
    if !isPositive {
        return -res
    } else {
        return res
    }
}

