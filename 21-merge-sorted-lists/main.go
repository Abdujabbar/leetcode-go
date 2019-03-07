package main

import "fmt"

func main() {
	l1 := ListNode{
		Val:  2,
		Next: nil,
	}
	// l1.Next = &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }
	// l1.Next.Next = &ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	l2 := ListNode{
		Val:  1,
		Next: nil,
	}
	// l2.Next = &ListNode{
	// 	Val:  3,
	// 	Next: nil,
	// }
	// l2.Next.Next = &ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	l3 := mergeTwoLists(&l1, &l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
