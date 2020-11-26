package linkList

import "testing"

func TestLinkList(t *testing.T) {
	node := New(1)
	node.Insert(2)
	node.Print()
}
func BenchmarkLinkList(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
