// Given a sorted array, two integers k and x, find the k closest elements to
// x in the array. The result should also be sorted in ascending order. If
// there is a tie, the smaller elements are always preferred.

func findClosestElements(arr []int, k int, x int) []int {
    i := sort.Search(len(arr) - k, func(i int) bool { return x - arr[i] <= arr[i+k] - x })
    return arr[i:i+k]
}
