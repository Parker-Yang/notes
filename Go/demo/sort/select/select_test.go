package _select

import "testing"

func TestSelect(t *testing.T) {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	Select(list)
	t.Log(list)
}
