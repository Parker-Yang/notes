package Stack

import "sync"

type Stack struct {
	Array []string
	Size  int
	Lock  sync.Mutex
}

// 入栈
func (s *Stack) Push(data string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Array = append(s.Array, data)
	s.Size++
}

// 出栈
func (s *Stack) Pop() string {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.Size == 0 {
		panic("stack is empty")
	}
	v := s.Array[s.Size-1]
	s.Array = s.Array[:s.Size-1]
	s.Size--
	return v
}

// 获取栈顶元素
func (s *Stack) Peek() string {
	if s.Size == 0 {
		panic("stack is empty")
	}
	v := s.Array[s.Size-1]
	return v
}

// 获取栈的长度
func (s *Stack) Len() int {
	return s.Size
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	if s.Size != 0 {
		return false
	} else {
		return true
	}
}
