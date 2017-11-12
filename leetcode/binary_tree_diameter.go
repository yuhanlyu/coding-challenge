// Given a binary tree, you need to compute the length of the diameter of the
// tree. The diameter of a binary tree is the length of the longest path
// between any two nodes in a tree. This path may or may not pass through the
// root.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    diameter, depth := helper(root.Left)
    rightDiameter, rightDepth := helper(root.Right)
    if rightDiameter > diameter {
        diameter = rightDiameter
    }
    if depth + rightDepth > diameter {
        diameter = depth + rightDepth
    }
    if depth < rightDepth {
        depth = rightDepth
    }
    return diameter, depth + 1
    
}

func diameterOfBinaryTree(root *TreeNode) int {
    diameter, _ := helper(root)
    return diameter
}
