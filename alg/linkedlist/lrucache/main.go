package main

import (
	"fmt"

	"github.com/nghiant3223/gopractice/alg/linkedlist/ds"
)

type LRUCache struct {
	count     int
	capacity  int
	head      *ds.DoubleNode
	tail      *ds.DoubleNode
	searchMap map[int]*ds.DoubleNode
}

func New(capacity int) LRUCache {
	c := LRUCache{}
	c.capacity = capacity
	c.searchMap = make(map[int]*ds.DoubleNode, c.capacity)
	return c
}

func (c *LRUCache) remove(n *ds.DoubleNode) {
	if c.head == n {
		c.head = n.Next
	}
	if c.tail == n {
		c.tail = n.Prev
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
}

func (c *LRUCache) append(n *ds.DoubleNode) {
	if c.head == nil {
		c.head = n
		c.tail = n
		return
	}
	n.Prev = c.tail
	c.tail.Next = n
	c.tail = n
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.searchMap[key]
	if !ok {
		return -1
	}
	c.remove(node)
	c.append(node)
	return node.Value
}

func (c *LRUCache) Put(key int, value int) {
	c.count++
	node, ok := c.searchMap[key]
	if ok {
		c.remove(node)
		c.append(node)
		return
	}
	newNode := &ds.DoubleNode{
		Key:   key,
		Value: value,
	}
	c.append(newNode)
	c.searchMap[key] = newNode
	if c.count > c.capacity {
		delete(c.searchMap, c.head.Key)
		c.remove(c.head)
		return
	}
}

func main() {
	capa := 2
	cache := New(capa)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
