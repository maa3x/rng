package rng

type intType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type numericType interface {
	intType | ~float32 | ~float64
}

func uniqueSlice[T ~[]E, E comparable](in T) T {
	seen := make(map[E]struct{})
	out := make(T, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func zeroVal[T any]() T {
	var zero T
	return zero
}
