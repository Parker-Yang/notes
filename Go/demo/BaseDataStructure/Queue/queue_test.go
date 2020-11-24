package Queue

import "testing"

func TestQueue(t *testing.T) {
	queue :=new(Queue)
	queue.Add("123")
	queue.Add("234")
	t.Log(queue.Remove())
	t.Log(queue.Remove())
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}