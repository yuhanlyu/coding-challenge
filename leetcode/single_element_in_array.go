// Given a sorted array consisting of only integers where every element appears
// twice except for one element which appears once. Find this single element
// that appears only once.

func singleNonDuplicate(nums []int) int {
    low, high := 0, len(nums) - 1
    for low < high {
        m := (low + high) / 2
        if nums[m] != nums[m ^ 1] {
            high = m
        } else {
            low = m + 1
        }
    }
    return nums[low]
}
