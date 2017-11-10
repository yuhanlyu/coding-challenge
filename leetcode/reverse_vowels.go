// Write a function that takes a string as input and reverse only the vowels of
// a string.

func reverseVowels(s string) string {
    result := []byte(s)
    for vowels, left, right := "aeiouAEIOU", 0, len(s) - 1; left < right ; {
        for strings.Index(vowels, string(result[left])) == -1 && left < right {
            left++
        }
        for strings.Index(vowels, string(result[right])) == -1 && left < right {
            right--
        }
        result[left], result[right], left, right = result[right], result[left], left + 1, right - 1
    }
    return string(result);
}
