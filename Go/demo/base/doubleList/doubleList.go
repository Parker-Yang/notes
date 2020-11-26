package doubleList

import (
	"sync"
)

// 双向链表
type DoubleList struct {
	Head *Node
	Tail *Node
	Size int
	Lock sync.Mutex
}

// 节点
type Node struct {
	Pre   *Node
	Next  *Node
	Value interface{}
}
// 从头部添加值
func (d *DoubleList) AddNodeFormHead(n int, data interface{}) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	node := new(Node)
	node.Value = data
	if d.Head.Next == nil {
		d.Head.Next = node
		node.Pre = d.Head
		d.Tail.Pre = node
		node.Next = d.Tail
	}
	if n > d.Size {
		panic("overSize")
	}
	dst := d.Head
	for i := 0; i < n; i++ {
		dst = dst.Next
	}
	node.Next = dst.Next
	node.Pre = dst
	dst.Next = node
	dst.Next.Pre = node
	d.Size++
}
