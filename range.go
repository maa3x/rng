package rng

// Range generates a random number of type T within the specified range [min, max].
// It panics if min < 0 or max <= 0.
func Range[T numericType](min, max T) T {
	if min < 0 {
		panic("min must be greater than or equal to 0")
	}
	if max <= 0 {
		panic("max must be greater than 0")
	}
	if min >= max {
		return max
	}

	switch any(min).(type) {
	case int, int8, int16, int32, int64:
		return T(r.Int64N(int64(max-min))) + min
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return T(r.Uint64N(uint64(max-min))) + min
	case float32, float64:
		return T(r.Float64()*(float64(max-min))) + min
	default:
		panic("unsupported type")
	}
}
