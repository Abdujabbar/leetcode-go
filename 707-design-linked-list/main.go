package main

import "fmt"

func main() {
	ll := Constructor()
	fmt.Println(ll.Get(0))
	ll.AddAtIndex(1, 2)
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))
	ll.AddAtIndex(0, 1)
	fmt.Println(ll.Get(0))
	fmt.Println(ll.Get(1))
}
