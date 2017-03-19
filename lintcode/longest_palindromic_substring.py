class Solution:
    # @param {string} s input string
    # @return {string} the longest palindromic substring
    def longestPalindrome(self, s):
        T = '#'.join('^{}$'.format(s))
        P, C, R, max, center = [0] * len(T), 0, 0, 0, 0
        for i in xrange(1, len(T) - 1):
            P[i] = min(R - i, P[2 * C - i]) if R >= i else 1
            while T[i + 1 + P[i]] == T[i - 1 - P[i]]:
                P[i] += 1
            if P[i] > max: max, center = P[i], i
            if i + P[i] > R:
                C, R = i, i + P[i]
        return s[(center  - max) / 2: (center  + max) / 2]
