// Given two integers x and y, calculate the Hamming distance.

func hammingDistance(x int, y int) int {
    z := x ^ y
    z = (z & 0x55555555) + ((z >> 1) & 0x55555555)
    z = (z & 0x33333333) + ((z >> 2) & 0x33333333)
    z = (z & 0x0f0f0f0f) + ((z >> 4) & 0x0f0f0f0f)
    z += (z >> 8)
    z += (z >> 16)
    return z & 0x3f
}
