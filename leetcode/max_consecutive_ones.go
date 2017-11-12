// Given a binary array, find the maximum number of consecutive 1s in this
// array.

func findMaxConsecutiveOnes(nums []int) int {
    current, max := 0, 0
    for _, num := range nums {
        if num == 1 {
            current++
        } else {
            current = 0
        }
        if current > max {
            max = current
        }
    }
    return max
}
