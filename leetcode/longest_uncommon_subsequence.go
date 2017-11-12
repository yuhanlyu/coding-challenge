// Given a group of two strings, you need to find the longest uncommon
// subsequence of this group of two strings. The longest uncommon subsequence
// is defined as the longest subsequence of one of these strings and this
// subsequence should not be any subsequence of the other strings.

func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    if len(a) > len(b) {
        return len(a)
    }
    return len(b)
}
