// Given a non negative integer number num. For every numbers i in the range
// 0 ≤ i ≤ num calculate the number of 1's in their binary representation and
// return them as an array.

func countBits(num int) []int {
    dp := make([]int, num + 1)
    for i := 1; i <= num; i++ {
        dp[i] = dp[i >> 1] + (i & 1)
    }
    return dp
}
