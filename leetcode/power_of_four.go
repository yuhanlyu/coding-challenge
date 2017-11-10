// Given an integer (signed 32 bits), write a function to check whether it is
// a power of 4.

func isPowerOfFour(num int) bool {
    return num > 0 && (num&(num-1)) == 0 && (num & 0x55555555) != 0
}
