// Determine the maximum amount of money the thief can rob tonight without
// alerting the police.

func rob(root *TreeNode) int {
    steal, notSteal := helper(root)
    if steal > notSteal {
        return steal
    }
    return notSteal
}

func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    stealLeft, notStealLeft := helper(root.Left)
    stealRight, notStealRight := helper(root.Right)
    maxLeft, maxRight := stealLeft, stealRight
    if stealLeft < notStealLeft {
        maxLeft = notStealLeft
    }
    if stealRight < notStealRight {
        maxRight = notStealRight
    }
    return root.Val + notStealLeft + notStealRight, maxLeft + maxRight
}
