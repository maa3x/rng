package main

import (
	"fmt"
	"slices"

	"github.com/maa3x/rng"
)

func main() {
	n := 10000
	picks := make([]float32, n)
	for i := range n {
		picks[i] = rng.Range[float32](0, 3)
	}
	slices.Sort(picks)
	fmt.Println(picks)
}
