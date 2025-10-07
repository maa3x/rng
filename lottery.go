package rng

import (
	"sync"
)

type lotteryItem[T any] struct {
	Value     T
	Weight    float64
	DrawCount int
}

// Lottery is a thread-safe generic lottery system that allows weighted random selection
// of items of type T. It maintains a collection of lottery items, each with an associated
// weight that determines the probability of selection.
type Lottery[T any] struct {
	mu    sync.Mutex
	items []*lotteryItem[T]
}

// NewLottery creates a new Lottery instance with the provided initial items.
func NewLottery[T any](items ...T) *Lottery[T] {
	return (&Lottery[T]{}).Append(items...)
}

// Append adds one or more values to the lottery with a default weight of 1.
func (l *Lottery[T]) Append(values ...T) *Lottery[T] {
	l.mu.Lock()
	for i := range values {
		l.items = append(l.items, &lotteryItem[T]{Weight: 1, Value: values[i]})
	}
	l.mu.Unlock()

	return l
}

// AppendWithWeight adds one or more values to the lottery with the specified weight.
// The weight determines the probability of each value being selected during a draw.
// Higher weights increase the likelihood of selection. weight <= 0 means the item will never be drawn.
func (l *Lottery[T]) AppendWithWeight(weight float64, values ...T) *Lottery[T] {
	l.mu.Lock()
	for i := range values {
		l.items = append(l.items, &lotteryItem[T]{Weight: weight, Value: values[i]})
	}
	l.mu.Unlock()

	return l
}

// Draw selects and returns a random item from the lottery based on weighted probabilities.
//
// If the lottery is empty, it returns the zero value of type T.
// Items with weight <= 0 are excluded from the weighted selection but may still
// be returned as a fallback if no weighted items exist.
//
// The selection algorithm uses cumulative weight distribution for fair randomization.
func (l *Lottery[T]) Draw() T {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.items) == 0 {
		return zeroVal[T]()
	}

	totalWeight := 0.0
	for i := range l.items {
		if l.items[i].Weight > 0 {
			totalWeight += l.items[i].Weight
		}
	}

	v := r.Float64() * totalWeight // [0, totalWeight]
	cumulative := 0.0
	for i := range l.items {
		if l.items[i].Weight <= 0 {
			continue
		}

		cumulative += l.items[i].Weight
		if v < cumulative {
			l.items[i].DrawCount++
			return l.items[i].Value
		}
	}

	i := r.IntN(len(l.items))
	l.items[i].DrawCount++
	return l.items[i].Value
}

func (l *Lottery[T]) DrawN(n int) []T {
	out := make([]T, n)
	for i := range n {
		out[i] = l.Draw()
	}
	return out
}

// Clear removes all items from the lottery, making it empty.
func (l *Lottery[T]) Clear() {
	l.mu.Lock()
	l.items = nil
	l.mu.Unlock()
}

// Size returns the number of items currently in the lottery.
func (l *Lottery[T]) Size() int {
	l.mu.Lock()
	size := len(l.items)
	l.mu.Unlock()
	return size
}

// Items returns a copy of all lottery items with their associated weights and values.
func (l *Lottery[T]) Items() []lotteryItem[T] {
	l.mu.Lock()
	clone := make([]lotteryItem[T], len(l.items))
	for i := range l.items {
		clone[i] = *l.items[i]
	}
	l.mu.Unlock()

	return clone
}
