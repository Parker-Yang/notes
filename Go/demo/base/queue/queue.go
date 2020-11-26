package queue

import "sync"

type Node struct {
	Next  *Node
	Value interface{}
}

type Queue struct {
	Node *Node
	Size int
	Lock sync.Mutex
}

func (q *Queue) Add(data interface{}) {
	// 线程安全
	q.Lock.Lock()
	defer q.Lock.Unlock()
	n := new(Node)
	n.Value = data
	if q.Node == nil {
		q.Node = n
	} else {
		for q.Node.Next != nil {
			q.Node = q.Node.Next
		}
		q.Node.Next = n
	}
	q.Size++
}

func (q *Queue) Remove() interface{} {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	if q.Node == nil {
		panic("queue is empty")
	}
	v := q.Node.Value
	q.Node = q.Node.Next
	q.Size--
	return v
}
