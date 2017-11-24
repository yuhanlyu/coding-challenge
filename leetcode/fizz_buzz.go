// Write a program that outputs the string representation of numbers from 1 to
// n.
// But for multiples of three it should output “Fizz” instead of the number
// and for the multiples of five output “Buzz”. For numbers which are
// multiples of both three and five output “FizzBuzz”.

func fizzBuzz(n int) []string {
    result := []string{}
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            result = append(result, "FizzBuzz")
        } else if i % 3 == 0 {
            result = append(result, "Fizz")
        } else if i % 5 == 0 {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(i))
        }
    }
    return result
}
