package leetcode

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var m [][]int
	if root == nil {
		return m
	}
	var currentLevelNodeList = []*TreeNode{root}

	for true {
		if len(currentLevelNodeList) <= 0 {
			break
		}
		var currentList []int
		var nextLevelNodeList []*TreeNode
		for _, node := range currentLevelNodeList {
			currentList = append(currentList, node.Val)
			if node.Left != nil {
				nextLevelNodeList = append(nextLevelNodeList, node.Left)
			}
			if node.Right != nil {
				nextLevelNodeList = append(nextLevelNodeList, node.Right)
			}
		}
		m = append(m, currentList)
		currentLevelNodeList = nextLevelNodeList
	}
	return m
}
