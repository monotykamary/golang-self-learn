package main

import (
	"fmt"
	"math"
)

/*
	Create a tree node that can refer to other tree nodes as left or right.
*/
type TreeNode struct {
	value float64
	left  *TreeNode
	right *TreeNode
}

/*
	Create a factory function to create our TreeNodes.
*/
func newTreeNode(value float64, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{value, left, right}
}

/*
	Create an object to store our maxPath data for each tree we calculate.
*/
type TreeData struct {
	maxPath float64
}

/*
	Calculate the max path sum by using using our post order function to mutate
	the maxPath variable.
*/
func (t *TreeData) MaxPathSum(root *TreeNode) float64 {
	t.PostOrder(root)
	return t.maxPath
}

/*
	Create a post order recursive function to check the maximum length as we
	traverse the tree recursively.
*/
func (t *TreeData) PostOrder(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	left := math.Max(t.PostOrder(root.left), 0)
	right := math.Max(t.PostOrder(root.right), 0)

	t.maxPath = math.Max(t.maxPath, left+right+root.value)
	return math.Max(left, right) + root.value
}

/*
	Create a utility function to print our tree recursively.
*/
func (t *TreeData) PrintRoot(root *TreeNode) interface{} {
	if root == nil {
		return nil
	}

	value := root.value
	left := t.PrintRoot(root.left)
	right := t.PrintRoot(root.right)

	rootSlice := []interface{}{value, left, right}
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

	tree1Data := TreeData{}
	tree2Data := TreeData{}

	fmt.Printf("Input: root = %v\n", tree1Data.PrintRoot(root1))
	fmt.Printf("Output: %v\n", tree1Data.MaxPathSum(root1))

	fmt.Printf("Input: root = %v\n", tree2Data.PrintRoot(root2))
	fmt.Printf("Output: %v\n", tree2Data.MaxPathSum(root2))
}
