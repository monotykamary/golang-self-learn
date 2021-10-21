package main

import (
	"fmt"
	"math"
)

var max float64

type TreeNode struct {
	value float64
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(value float64, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{value, left, right}
}

func maxPathSum(root *TreeNode) float64 {
	max = math.Inf(-1)
	postOrder(root)
	return max
}

func postOrder(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	left := math.Max(postOrder(root.left), 0)
	right := math.Max(postOrder(root.right), 0)

	max = math.Max(max, left+right+root.value)
	return math.Max(left, right) + root.value
}

func printRoot(root *TreeNode) interface{} {
	if root == nil {
		return nil
	}

	value := root.value
	left := printRoot(root.left)
	right := printRoot(root.right)

	var rootSlice []interface{} = []interface{}{value, left, right}
	return rootSlice
}

func main() {
	root1 := newTreeNode(1, newTreeNode(2, nil, nil), newTreeNode(3, nil, nil))
	root2 := newTreeNode(
		-10,
		newTreeNode(9, nil, nil),
		newTreeNode(20,
			newTreeNode(15, nil, nil),
			newTreeNode(7, nil, nil),
		),
	)

	fmt.Printf("Input: root = %v\n", printRoot(root1))
	fmt.Printf("Output: %v\n", maxPathSum(root1))

	fmt.Printf("Input: root = %v\n", printRoot(root2))
	fmt.Printf("Output: %v\n", maxPathSum(root2))
}
