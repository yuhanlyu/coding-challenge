// Assume you are an awesome parent and want to give your children some
// cookies. But, you should give each child at most one cookie. Each child i
// has a greed factor gi, which is the minimum size of a cookie that the child
// will be content with; and each cookie j has a size sj. If sj >= gi, we can
// assign the cookie j to the child i, and the child i will be content. Your
// goal is to maximize the number of your content children and output the
// maximum number.

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    result := 0
    for i := 0; result < len(g) && i < len(s); i++ {
        if g[result] <= s[i] {
            result++
        }
    }
    return result
}
