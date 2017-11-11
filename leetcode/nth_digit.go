// Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8,
// 9, 10, 11, ...

func findNthDigit(n int) int {
    length, count, start := 1, 9, 1
    for n > length * count {
        n -= length * count
        length, count, start = length + 1, count * 10, start * 10
    }
    target := strconv.Itoa(start + (n - 1) / length)
    num, _ := strconv.Atoi(string(target[(n - 1) % length]))
    return num
}
