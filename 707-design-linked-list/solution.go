package main

type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	Head *Node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
		size: 0,
	}
}

func (ll *MyLinkedList) incrementSize() {
	ll.size++
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (ll *MyLinkedList) Get(index int) int {
	c := ll.Head
	for {
		if c == nil {
			return -1
		}
		if index == 0 && c != nil {
			return c.Val
		}
		if index <= 0 {
			break
		}
		index--
		c = c.Next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (ll *MyLinkedList) AddAtHead(val int) {
	if ll.Head == nil {
		ll.Head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		c := ll.Head
		ll.Head = &Node{
			Val:  val,
			Next: c,
		}
	}
	ll.incrementSize()
}

/** Append a node of value val to the last element of the linked list. */
func (ll *MyLinkedList) AddAtTail(val int) {
	if ll.Head == nil {
		ll.Head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		c := ll.Head
		for c.Next != nil {
			c = c.Next
		}
		c.Next = &Node{
			Val:  val,
			Next: nil,
		}
	}
	ll.incrementSize()
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > ll.size {
		return
	}
	if ll.Head == nil {
		ll.Head = &Node{
			Val:  val,
			Next: nil,
		}
	} else {
		c := ll.Head
		var prev *Node
		for {
			if index == 0 {
				vv := Node{
					Val:  val,
					Next: c,
				}
				if prev != nil {
					prev.Next = &vv
				} else {
					ll.Head = &Node{
						Val:  val,
						Next: c,
					}
				}
				break
			}
			if index <= 0 {
				break
			}
			prev = c
			c = c.Next
			if c == nil {
				prev.Next = &Node{
					Val:  val,
					Next: nil,
				}
				return
			}
			index--
		}
	}
	ll.incrementSize()
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	c := ll.Head
	var prev *Node
	for {
		if index == 0 {
			if prev == nil {
				ll.Head = c.Next
			} else {
				prev.Next = c.Next
			}
			ll.size--
			break
		}
		prev = c
		c = c.Next
		if c == nil {
			break
		}
		index--
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
