// We define a harmonious array is an array where the difference between its
// maximum value and its minimum value is exactly 1.
// Now, given an integer array, you need to find the length of its longest
// harmonious subsequence among all its possible subsequences.

func findLHS(nums []int) int {
    result, table := 0, map[int]int{}
    for _, num := range nums {
        table[num]++
    }
    for key, count1 := range table {
        if count2, ok := table[key + 1]; ok && count1 + count2 > result {
            result = count1 + count2
        }
    }
    return result
}
