// Given an integer array, you need to find one continuous subarray that if
// you only sort this subarray in ascending order, then the whole array will
// be sorted in ascending order, too.

func findUnsortedSubarray(nums []int) int {
    begin, end := -1, -2
    for min, max, i := nums[len(nums) - 1], nums[0], 1; i < len(nums); i++ {
        if nums[i] >= max {
            max = nums[i]
        } else {
            end = i;
        }
        if nums[len(nums) - 1 - i] <= min {
            min = nums[len(nums) - 1 - i]
        } else {
            begin = len(nums) - 1 - i
        }
    }
    return end - begin + 1
}
