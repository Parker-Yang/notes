package LinkList

import "fmt"

// 节点（多个节点构成一条链表）
type Node struct {
	Next *Node
	Data interface{}
}

func New(data interface{}) *Node {
	return &Node{
		Next: nil,
		Data: data,
	}
}

func (n *Node) Insert(data interface{}) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = New(data)
}

func (n *Node) Print() {
	for n != nil {
		fmt.Println(n.Data)
		n = n.Next
	}
}
