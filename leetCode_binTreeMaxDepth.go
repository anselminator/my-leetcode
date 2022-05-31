package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	c := TreeNode{1, nil, nil}
	b := TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	e := TreeNode{5, &a, &b}
	f := TreeNode{0, &TreeNode{42, &c, &c}, &TreeNode{99, &a, &e}}
	g := TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println("maxdepth")
	fmt.Println(maxDepth(&a))
	fmt.Println(maxDepth(&b))
	fmt.Println(maxDepth(&c))
	fmt.Println(maxDepth(&e))
	fmt.Println(maxDepth(&f))
	fmt.Println(maxDepth(&g))

}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	} else {
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		return 1 + max(l, r)
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
