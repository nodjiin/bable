package buffers

import "testing"

func TestPut(t *testing.T) {
	cb := NewCircBuffer(2)

	cb.Put(1)
	cb.Put(2)
	cb.Put(3)

	if cb.buffer[0] != 3 {
		t.Error("CircBuffer.Travers: array[0] != 3")
	}

	if cb.buffer[1] != 2 {
		t.Error("CircBuffer.Travers: array[1] != 2")
	}

}

func TestTraverse(t *testing.T) {
	cb := NewCircBuffer(2)

	cb.Put(1)
	cb.Put(2)

	var array [2]int
	index := 0
	op := func(value interface{}) {
		array[index] = value.(int)
		index++
	}

	cb.Traverse(op)

	if array[0] != 1 {
		t.Error("CircBuffer.Travers: array[0] != 1")
	}

	if array[1] != 2 {
		t.Error("CircBuffer.Travers: array[1] != 2")
	}
}
