/**
 * @param a: A 32bit integer
 * @param b: A 32bit integer
 * @param n: A 32bit integer
 * @return: An integer
 */
func fastPower (a int, b int, n int) int {
    if b == 1 {
        return 0
    }
    cur := 1
    var curMod int64 = int64(a) % int64(b)
    var res int64 = 1
    for i := 0; i < 32; i++ {
        if n & cur > 0 {
            res = (res * curMod) % int64(b)
        }
        cur = cur << 1
        curMod = (curMod * curMod) % int64(b)
    }
    return int(res)
}

