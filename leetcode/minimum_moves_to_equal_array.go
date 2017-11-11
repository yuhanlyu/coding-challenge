// Given a non-empty integer array of size n, find the minimum number of moves
// required to make all array elements equal, where a move is incrementing
// n - 1 elements by 1.
func minMoves(nums []int) int {
    min, result := nums[0], 0
    for _, num := range nums {
        if num < min {
            min = num
        }
    }
    for _, num := range nums {
        result += num - min
    }
    return result
}
