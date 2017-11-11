// Given a string, find the first non-repeating character in it and return
// it's index. If it doesn't exist, return -1.

func firstUniqChar(s string) int {
    letters := map[rune]int{}
    for _, letter := range s {
        letters[letter]++
    }
    for index, letter := range s {
        if letters[letter] == 1 {
            return index
        }
    }
    return -1
}
