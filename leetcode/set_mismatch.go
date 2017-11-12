// The set S originally contains numbers from 1 to n. But unfortunately, due
// to the data error, one of the numbers in the set got duplicated to another
// number in the set, which results in repetition of one number and loss of
// another number.
// Given an array nums representing the data status of this set after the
// error. Your task is to firstly find the number occurs twice and then find
// the number that is missing. Return them in the form of an array.

func findErrorNums(nums []int) []int {
    c, countBit, countOriginal := 0, 0, 0
    for i, num := range nums {
        c ^= num ^ (i + 1)
    }
    mask, result := c ^ (c & (c - 1)), 0
    for i, num := range nums {
        if num & mask > 0 {
            countBit++
            result ^= num
        }
        if ((i + 1) & mask) > 0 {
            countOriginal++
            result ^= (i + 1)
        }
    }
    if countBit > countOriginal {
        return []int{result, c ^ result}
    }
    return []int{c ^ result, result}
}
