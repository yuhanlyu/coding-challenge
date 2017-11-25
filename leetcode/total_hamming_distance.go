// Find the total Hamming distance between all pairs of the given numbers.

func totalHammingDistance(nums []int) int {
    result := 0
    for i := uint(0); i < 32; i++ {
        count := 0
        for _, num := range nums {
            count += (num >> i) & 1
        }
        result += count * (len(nums) - count)
    }
    return result
}
