// Given an arbitrary ransom note string and another string containing letters
// from all the magazines, write a function that will return true if the
// ransom note can be constructed from the magazines ; otherwise, it will
// return false.
// Each letter in the magazine string can only be used once in your ransom note.

func canConstruct(ransomNote string, magazine string) bool {
    letters := map[rune]int{}
    for _, letter := range magazine {
        letters[letter]++
    }
    for _, letter := range ransomNote {
        if letters[letter] == 0 {
            return false
        }
        letters[letter]--
    }
    return true
}
