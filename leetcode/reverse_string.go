// Write a function that takes a string as input and returns the string
// reversed.

func reverseString(s string) string {
    reverse := make([]byte, len(s), len(s))
    for i, c := range []byte(s) {
        reverse[len(s) - i - 1] = c
    }
    return string(reverse[:])
}
