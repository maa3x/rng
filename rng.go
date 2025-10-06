package rng

import (
	"math/rand/v2"
)

var r = rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))

// ReplaceRandSource replaces the global random number generator's source with the provided source.
//
// This function allows for customizing the randomization behavior by using a different
// source implementation, such as for testing with deterministic seeds or using
// alternative random number generation algorithms.
func ReplaceRandSource(src rand.Source) {
	r = rand.New(src)
}

// Num generates a random number of the specified numeric type.
//
// For floating-point types (float32, float64), it returns a random value between 0.0 and 1.0.
// For integer types, it returns a random value within the full range of that type.
func Num[Number numericType]() Number {
	switch any(zeroVal[Number]()).(type) {
	case float32, float64:
		return Number(r.Float64())
	default:
		return Number(r.Uint64())
	}
}

// N returns a non-negative pseudo-random number in the range [0, n]. It panics if n <= 0.
func N[Int intType](n Int) Int {
	if n < 0 {
		panic("n must be non-negative")
	}
	return Int(r.Uint64N(uint64(n)))
}
