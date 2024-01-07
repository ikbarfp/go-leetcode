package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func anotherSortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	median := len(nums) / 2
	root := &TreeNode{Val: nums[median]}
	root.Left = sortedArrayToBST(nums[:median])
	root.Right = sortedArrayToBST(nums[median+1:])
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	median := len(nums) / 2
	root := &TreeNode{Val: nums[median]}
	if median != 0 {
		root.Left = sortedArrayToBST(nums[:median])
	}
	if median != len(nums)-1 {
		root.Right = sortedArrayToBST(nums[median+1:])
	}
	return root
}
