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
	fmt.Println("walking the tree inorder (left,node,right)")
	fmt.Println(inorderTraversal(&a))
	fmt.Println(inorderTraversal(&b))
	fmt.Println(inorderTraversal(&c))
	fmt.Println(inorderTraversal(&e))
	fmt.Println(inorderTraversal(&f))

}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	} else {
		//		l := inorderTraversal(root.Left)
		//		r := inorderTraversal(root.Right)
		//		res = append(res, l...)
		//		res = append(res, root.Val)
		//		res = append(res, r...)
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)

		return res
	}
}
