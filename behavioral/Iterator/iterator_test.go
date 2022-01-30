package Iterator

import "testing"
import "fmt"

func TestArrayIterator(t *testing.T) {
	array := []interface{}{1,3,9,2,8,7}
	index := 0
	iterator := ArrayIterator{array: array, index: &index}
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		for value != array[*index] {
			fmt.Println("error")
		}
		fmt.Println(*index, value)
	}

}