package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	l2 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	l3 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	l4 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	l5 := ListNode{1, &ListNode{2, nil}}
	l6 := ListNode{1, &ListNode{2, nil}}
	l6.Next.Next = &l6
	l7 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	l7.Next.Next.Next.Next.Next.Next = l7.Next.Next

	l8 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}}

	fmt.Println("l1 is a Palindrome:", isPalindrome(&l1))
	fmt.Println("l8 is a Palindrome:", isPalindrome(&l8))

	fmt.Println("l2 has cycles:", hasCycle(&l2))
	fmt.Println("l3 has cycles:", hasCycle(&l3))
	fmt.Println("l4 has cycles:", hasCycle(&l4))
	fmt.Println("l6 has cycles:", hasCycle(&l6))
	fmt.Println("l7 has cycles:", hasCycle(&l7))

	displayList(&l1)
	fmt.Printf("\n")
	fmt.Println(" , with 2. from last removed:", removeNthFromEnd(&l1, 2))
	displayList(&l1)
	fmt.Printf("\n")

	fmt.Println(" , with last removed:", removeNthFromEnd(&l2, 1))
	displayList(&l2)
	fmt.Printf("\n")

	fmt.Println(" , with 5. from last removed:", removeNthFromEnd(&l4, 5))
	displayList(&l4)
	fmt.Printf("\n")

	fmt.Println(" , with 1. from last removed:", removeNthFromEnd(&l5, 1))
	displayList(&l5)
	fmt.Printf("\n")

	fmt.Println(" , with 1. from last removed:", removeNthFromEnd(&l5, 1))
	displayList(&l5)
	fmt.Printf("\n")

	fmt.Println(" , with 7. from last removed:", removeNthFromEnd(&l3, 7))
	displayList(&l3)
	fmt.Printf("\n")

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := head
	b4rem := head
	if head.Next == nil {
		head = nil
		return nil
	}
	for i := 0; i < n; i++ {
		tmp = tmp.Next
	}
	if tmp == nil {
		head = head.Next
		return head
	}
	b4rem = head
	//	displayList(head)
	//	fmt.Println(" \n b4rem: ")

	for tmp.Next != nil {
		tmp = tmp.Next
		b4rem = b4rem.Next
	}
	//	displayList(head)
	//	fmt.Println(" \n b4rem: ")
	//	displayList(b4rem)
	if b4rem.Next == nil {
		return nil
	}
	b4rem.Next = b4rem.Next.Next
	//	fmt.Println(" \n b4rem: ")
	//	displayList(b4rem)

	return head
}

func displayList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
}

func hasCycle(head *ListNode) bool {
	var visited map[*ListNode]bool = make(map[*ListNode]bool)
	if head == nil {
		return false
	}
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}

func isPalindrome(head *ListNode) bool {
	var isPalin *ListNode = &ListNode{}
	tmp := head
	for tmp != nil {
		insertBefore(isPalin, head.Val)
		tmp = tmp.Next
	}
	displayList(head)
	displayList(isPalin)

	for head.Val == isPalin.Val {
		head = head.Next
		isPalin = isPalin.Next
	}
	return head == nil && isPalin == nil || isPalin.Val == head.Val
}

func insertBefore(head *ListNode, el int) {
	if head == nil {
		head = &ListNode{el, nil}
	} else {
		new := &ListNode{el, head}
		head = new
	}
}
