package doubleList

import (
	"fmt"
	"testing"
)

func TestDoubleList(t *testing.T) {
	list := new(DoubleList)
	list.AddNodeFormHead(0,1)
	fmt.Println(list)
}
