// Given a string which consists of lowercase or uppercase letters, find the
// length of the longest palindromes that can be built with those letters.

func longestPalindrome(s string) int {
    letters, count := map[rune]bool{}, 0
    for _, letter := range s {
        if _, ok := letters[letter]; ok{
            count++
            delete(letters, letter)
        } else {
            letters[letter] = true
        }
    }
    if len(letters) != 0 {
        return 2 * count + 1
    }
    return 2 * count
}
