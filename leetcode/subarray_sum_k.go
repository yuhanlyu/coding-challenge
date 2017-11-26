// Given an array of integers and an integer k, you need to find the total
// number of continuous subarrays whose sum equals to k.

func subarraySum(nums []int, k int) int {
    result, sum, table := 0, 0, map[int]int{0: 1}
    for _, num := range nums {
        sum += num
        if prefix, ok := table[sum - k]; ok {
            result += prefix
        }
        table[sum]++
    }
    return result
}
