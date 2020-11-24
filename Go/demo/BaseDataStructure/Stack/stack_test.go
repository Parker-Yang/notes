package Stack

import "testing"

func TestStack(t *testing.T) {
	arr := new(Stack)
	t.Log(arr.IsEmpty())
	arr.Push("s")
	arr.Push("t")
	arr.Push("r")
	arr.Push("i")
	arr.Push("g")
	t.Log(arr.Len())
	arr.Pop()
	t.Log(arr.Peek())
}
func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}