// Calculate the sum of two integers a and b, but you are not allowed to use
// the operator + and -.

func getSum(a int, b int) int {
    if a == 0 {
        return b
    }
    if b == 0 {
        return a
    }

    for b != 0 {
        a, b = a ^ b, (a & b) << 1
    }
	return a;
}
