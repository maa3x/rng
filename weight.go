package rng

// Weight represents a non-negative weight value.
type Weight uint

// WeightMap represents a mapping from comparable keys to their associated weights.
// It is commonly used in random selection algorithms where each key
// has a corresponding weight that influences its probability of being selected.
type WeightMap[K comparable] map[K]Weight

// Pick selects a random key from the WeightMap based on the weights of the keys.
//
// Keys with higher weights have a proportionally higher chance of being selected.
// If the total weight is zero (all weights are zero or the map is empty),
// it returns the zero value of type K.
// The selection uses a cumulative distribution where each key's probability
// is its weight divided by the total weight of all keys.
func (wm WeightMap[K]) Pick() K {
	var total Weight
	for _, v := range wm {
		total += v
	}
	if total == 0 {
		return zeroVal[K]()
	}

	r := r.Uint64N(uint64(total))
	var cumulative Weight
	for k, v := range wm {
		cumulative += v
		if r < uint64(cumulative) {
			return k
		}
	}
	return zeroVal[K]()
}
