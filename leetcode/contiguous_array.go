// Given a binary array, find the maximum length of a contiguous subarray with
// equal number of 0 and 1.

func findMaxLength(nums []int) int {
    result, count, table := 0, 0, map[int]int{0: 0,}
    for i, num := range nums {
        if num == 0 {
            count--
        } else {
            count++
        }
        if previous, ok := table[count]; ok {
            if i - previous + 1 > result {
                result = i - previous + 1
            }
        } else {
            table[count] = i + 1
        }
    }
    return result
}
