// Given two arrays, write a function to compute their intersection.

func intersection(nums1 []int, nums2 []int) []int {
    elements, result := map[int]bool{}, []int{}
    for _, num := range nums1 {
        elements[num] = true
    }
    for _, num := range nums2 {
        if elements[num] {
            result = append(result, num)
            elements[num] = false
        }
    }
    return result
}
