package buffers

type CircBuffer struct {
	buffer     []interface{}
	writeIndex uint
	size       uint
}

func NewCircBuffer(size uint) *CircBuffer {
	var buff CircBuffer
	buff.buffer = make([]interface{}, size)
	buff.size = size
	buff.writeIndex = 0
	return &buff
}

func (cb *CircBuffer) Put(item interface{}) {
	cb.buffer[cb.writeIndex] = item
	cb.writeIndex = (cb.writeIndex + 1) % cb.size
}

// Traverse the buffer and apply op to each value.
// The operation will be executed on all the element of the buffer,
// even if it's not completely full
func (cb *CircBuffer) Traverse(op func(interface{})) {
	op(cb.buffer[cb.writeIndex])
	readIndex := cb.writeIndex + 1
	for readIndex != cb.writeIndex {
		op(cb.buffer[readIndex])
		readIndex = (readIndex + 1) % cb.size
	}
}
