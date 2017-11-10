// Given two arrays, write a function to compute their intersection. 

func intersect(nums1 []int, nums2 []int) []int {
    elements, result := map[int]int{}, []int{}
    for _, num := range nums1 {
        elements[num]++
    }
    for _, num := range nums2 {
        if elements[num] > 0{
            result = append(result, num)
            elements[num]--
        }
    }
    return result
}
