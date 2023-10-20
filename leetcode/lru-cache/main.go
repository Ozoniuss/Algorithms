package main

import "fmt"

/*
This is the exact solution I started writing at about 1:30 AM and finished
writing at about 2:30 AM. If you think it looks ugly whatever. It beats ~93%
speed and ~88% memory so f off.
*/

type Node struct {
	value int
	key   int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	vals  map[int]*Node
	cap   int
	first *Node
	last  *Node
}

func Constructor(capacity int) LRUCache {
	vals := make(map[int]*Node, capacity)
	return LRUCache{
		vals:  vals,
		cap:   capacity,
		first: nil,
		last:  nil,
	}
}

func (this *LRUCache) Get(key int) int {

	if len(this.vals) == 0 {
		return -1
	}

	// not found
	var n *Node
	var ok bool
	if n, ok = this.vals[key]; !ok {
		return -1
	}

	ret := n.value

	// edge cases
	if len(this.vals) == 1 {
		return ret
	}
	if n == this.last {
		return ret
	}
	// swap
	n.next.prev = n.prev
	if n == this.first {
		this.first = n.next
	} else {
		n.prev.next = n.next
	}

	// set n as the last one, knowing it isnt already
	this.last.next = n
	n.prev = this.last
	this.last = n
	n.next = nil

	return ret
}

func (this *LRUCache) Put(key int, value int) {

	if len(this.vals) == 0 {
		n := &Node{
			value: value,
			key:   key,
			prev:  nil,
			next:  nil,
		}
		this.first = n
		this.last = n
		this.vals[key] = n
		return
	}

	// already exists, then it's essentially a get operation to change the
	// pointers but change the value as well
	if n, ok := this.vals[key]; ok {
		this.Get(key)
		n.value = value
		return
	}

	n := &Node{
		value: value,
		key:   key,
		prev:  this.last,
		next:  nil,
	}
	this.vals[key] = n
	this.last.next = n
	this.last = n

	// time to remove something
	if len(this.vals) > this.cap {
		if (this.cap) == 1 {
			// its easiest to just remake the map
			this.vals = make(map[int]*Node, this.cap)
			this.vals[key] = n
			this.first = n
			n.prev = nil
			n.next = nil
			return
		}
		k := this.first.key
		this.first.next.prev = nil
		this.first = this.first.next
		delete(this.vals, k)
	}
}

func main() {
	obj := Constructor(3)
	obj.Put(3, 4)
	fmt.Println(obj.Get(3))
	obj.Put(3, 5)
	fmt.Println(obj.Get(3))
	obj.Put(2, 7)
	obj.Put(4, 8)
	fmt.Println(obj.Get(2))
	obj.Put(1, 5)
	fmt.Println(obj.Get(3))

}
