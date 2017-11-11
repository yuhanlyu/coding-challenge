// Given a non-empty string check if it can be constructed by taking a
// substring of it and appending multiple copies of the substring together. 

func repeatedSubstringPattern(s string) bool {
    return strings.Contains((s + s)[1: 2 * len(s) - 1], s)
}
