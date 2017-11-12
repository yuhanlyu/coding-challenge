// Given a positive integer, check whether it has alternating bits: namely, if
// two adjacent bits will always have different values.

func hasAlternatingBits(n int) bool {
    n = n ^ (n >> 2)
    return (n & (n-1)) == 0
}
