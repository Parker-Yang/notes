package insert

import "testing"

func TestInsert(t *testing.T) {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	Insert(list)
	t.Log(list)
}