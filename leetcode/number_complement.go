// Given a positive integer, output its complement number. The complement
// strategy is to flip the bits of its binary representation.

func findComplement(num int) int {
    mask := num
    mask |= mask >> 1
    mask |= mask >> 2
    mask |= mask >> 4
    mask |= mask >> 8
    mask |= mask >> 16
    return num ^ mask
}
