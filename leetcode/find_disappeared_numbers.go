// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
// elements appear twice and others appear once.
// Find all the elements of [1, n] inclusive that do not appear in this array.

func findDisappearedNumbers(nums []int) []int {
    result, n := []int{}, len(nums)
    for i, _ := range nums {
        nums[(nums[i] - 1) % n] += n
    }
    for i, _ := range nums {
        if nums[i] <= n {
            result = append(result, i + 1)
        }
    }
    return result
}
