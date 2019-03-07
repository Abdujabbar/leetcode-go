package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

	l1.Next = mergeTwoLists(l1.Next, l2)

	return l1
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var t *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			if result == nil {
				result = &ListNode{
					Val:  l1.Val,
					Next: nil,
				}
				t = result
			} else {
				result.Next = &ListNode{
					Val:  l1.Val,
					Next: nil,
				}
				result = result.Next
			}
			l1 = l1.Next
		} else {
			if result == nil {
				result = &ListNode{
					Val:  l2.Val,
					Next: nil,
				}
				t = result
			} else {
				result.Next = &ListNode{
					Val:  l2.Val,
					Next: nil,
				}
				result = result.Next
			}
			l2 = l2.Next
		}
	}
	if l1 != nil {
		if result != nil {
			result.Next = l1
		} else {
			t = l1
		}
	}
	if l2 != nil {
		if result != nil {
			result.Next = l2
		} else {
			t = l2
		}
	}
	return t
}
