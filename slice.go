package rng

import (
	"fmt"
)

// Pick returns a random element from the provided slice.
// If the slice is empty, it returns the zero value of type E.
func Pick[E any](in []E) E {
	if len(in) == 0 {
		return zeroVal[E]()
	}
	return in[r.IntN(len(in))]
}

// PickN returns n randomly selected elements from the input slice.
//
// Elements may be duplicated if n is greater than the length of the input slice.
// Returns nil if n is less than or equal to 0 or if the input slice is empty.
// The selection uses a random permutation to determine which elements to pick.
func PickN[E any](in []E, n int) []E {
	if n <= 0 || len(in) == 0 {
		return nil
	}

	out := make([]E, n)
	p := r.Perm(n)
	ln := len(in)
	for i := range n {
		out[i] = in[p[i]%ln]
	}
	return out
}

// PickNDistinct returns n distinct elements randomly selected from the given slice.
//
// The function uses Fisher-Yates shuffle algorithm to ensure each element has an equal
// probability of being selected. The order of elements in the returned slice is randomized.
// Returns nil if n <= 0. Panics if n is greater than the length of the input slice.
func PickNDistinct[E any](slice []E, n int) []E {
	if n <= 0 {
		return nil
	}
	if n > len(slice) {
		panic("n must be less than the length of the slice")
	}

	perm := r.Perm(len(slice))
	result := make([]E, n)
	for i := range n {
		result[i] = slice[perm[i]]
	}

	return result
}

// PickNUnique returns n randomly selected unique elements from the given slice.
//
// It first removes duplicates from the input slice, then randomly picks n distinct elements.
// Returns an error if n is less than or equal to 0, or if n exceeds the number of unique elements in the slice.
func PickNUnique[S ~[]E, E comparable](slice S, n int) (S, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be greater than 0")
	}

	unique := uniqueSlice(slice)
	if n > len(unique) {
		return nil, fmt.Errorf("n must be less than or equal to the number of unique elements in the slice")
	}

	return PickNDistinct(unique, n), nil
}

// Shuffle randomly rearranges the elements of the slice in place using the Fisher-Yates shuffle algorithm.
//
// The function modifies the original slice and does not return a new slice.
// If the slice has 1 or fewer elements, no shuffling is performed.
func Shuffle[T any](in []T) {
	if len(in) <= 1 {
		return
	}
	r.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}
