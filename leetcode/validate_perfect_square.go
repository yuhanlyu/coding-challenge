// Given a positive integer num, write a function which returns True if num is
// a perfect square else False.

func isPerfectSquare(num int) bool {
    r := num
    for ; r * r > num; r = (r + num / r) / 2 {
    }
    return r * r == num
}
