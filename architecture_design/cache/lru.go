// see: https://medium.com/hackernoon/build-a-go-cache-in-10-minutes-c908a8255568
package main

import "fmt"

// CacheSize is the size of cache.
const CacheSize = 5

// Node is a node for queue.
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Queue is a double linked list.
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Hash is a maps string to node in queue.
type Hash map[string]*Node

// Cache is a LRU cache.
type Cache struct {
	Queue Queue
	Hash  Hash
}

// NewCache creates a new Cache.
func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

// NewQueue creates a new Queue.
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

// Check checks if data exists in cache.
func (c *Cache) Check(str string) string {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}

	c.Add(node)
	c.Hash[str] = node
	return node.Value
}

// Remove removes a node from Cache.
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Value)
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	c.Queue.Length--

	delete(c.Hash, n.Value)
	return n
}

// Add adds a node to Cache.
func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Value)
	tmp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > CacheSize {
		c.Remove(c.Queue.Tail.Left)
	}
}

// Display displays a Queue.
func (c *Cache) Display() {
	c.Queue.Display()
}

// Display displays a Queue.
func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	cache := NewCache()
	for _, word := range []string{"a", "b", "c", "d", "e", "a", "b", "c"} {
		cache.Check(word)
		cache.Display()
	}
}
