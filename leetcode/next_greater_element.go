// You are given two arrays (without duplicates) nums1 and nums2 where nums1’s
// elements are subset of nums2. Find all the next greater numbers for nums1's
// elements in the corresponding places of nums2.

func nextGreaterElement(findNums []int, nums []int) []int {
    nextGreaterElement, stack, top := map[int]int{}, map[int]int{}, 0
    for _, num := range nums {
        for top > 0 && stack[top - 1] < num {
            top--
            nextGreaterElement[stack[top]] = num
        }
        stack[top] = num
        top++
    }
    for i, num := range findNums {
        if ans, ok := nextGreaterElement[num]; ok {
            findNums[i] = ans
        } else {
            findNums[i] = -1
        }
    }
    return findNums;
}
