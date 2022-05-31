package main

import "fmt"

/*// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

func main() {
	var list1 *ListNode = &ListNode{0, nil}

	fmt.Println("list")
	fmt.Println(list1)
	insertBefore(list1, 99)
	insertBefore(list1, 42)
	fmt.Println(list1)

}

/*func isPalindrome(head *ListNode) bool {
	var head2 *ListNode = nil
	var head3 *ListNode = head
	for head3 != nil {
		insertBefore(head2, head3.Val)
		head3 = head3.Next
	}
	return listCompare(head, head3)
}*/

func listCompare(l1 *ListNode, l2 *ListNode) bool {
	if l1.Next == nil && l2.Next == nil {
		return true
	}
	if l1.Val != l2.Val {
		return false
	} else {
		return listCompare(l1.Next, l2.Next)
	}
}

/* func popList(head *ListNode) int {
	var v int = 0
	if head == nil {
		return nil
	} else {
		v = head.Val
		head = head.Next
		return v
	}
	head
}
*/
func insertBefore(head *ListNode, el int) {
	if head == nil {
		head = &ListNode{el, nil}
	} else {
		new := &ListNode{el, head}
		head = new
	}
}
