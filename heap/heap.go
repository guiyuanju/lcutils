package heap

type Heap[T any] struct {
	values []T
	less   func(i, j int) bool
}

func New[T any](values []T, less func(i, j int) bool) Heap[T] {
	return Heap[T]{values, less}
}

func (h Heap[T]) Len() int           { return len(h.values) }
func (h Heap[T]) Less(i, j int) bool { return h.less(i, j) }
func (h Heap[T]) Swap(i, j int)      { h.values[i], h.values[j] = h.values[j], h.values[i] }
func (h *Heap[T]) Push(x any)        { h.values = append(h.values, x.(T)) }
func (h *Heap[T]) Pop() any {
	old := h.values
	length := len(old)
	res := old[length-1]
	h.values = old[:length-1]
	return res
}
