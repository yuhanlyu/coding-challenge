// Given an array consisting of n integers, find the contiguous subarray of
// given length k that has the maximum average value. And you need to output
// the maximum average value.

func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    max := sum
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i - k]
        if sum > max {
            max = sum
        }
    }
    return float64(max) / float64(k)
}
