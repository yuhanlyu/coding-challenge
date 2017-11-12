// We define the Perfect Number is a positive integer that is equal to the sum
// of all its positive divisors except itself.
// Now, given an integer n, write a function that returns true when it is a
// perfect number and false when it is not.

func checkPerfectNumber(num int) bool {
    if num == 1 {
        return false
    }
    sum := 1
    for i := 2; i * i <= num; i++ {
        if num % i == 0 {
            sum += i + num / i
        }
    }
    return sum == num
}
