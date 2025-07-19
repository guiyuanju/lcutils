package math

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Numeric](xs []T) T {
	var res T
	for _, x := range xs {
		res += x
	}
	return res
}
