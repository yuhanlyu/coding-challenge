// Given a positive integer n, break it into the sum of at least two positive
// integers and maximize the product of those integers. Return the maximum
// product you can get.

func integerBreak(n int) int {
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    product := 1
    for ; n > 4; n -= 3 {
        product *= 3
    }
    return product * n
}
