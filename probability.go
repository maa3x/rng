package rng

var P50 = Probability(0.5)

// Probability represents a probability value between 0 and 1 (inclusive).
type Probability float64

func (p Probability) Check() bool {
	if p <= 0 {
		return false
	}
	if p >= 1 {
		return true
	}
	return r.Float64() < float64(p)
}

// ProbabilityMap is a map where keys have associated probabilities or weights.
type ProbabilityMap[K comparable] map[K]Probability

// Pick selects a key from the ProbabilityMap based on the associated probabilities or weights.
//
// Keys with higher values have a higher chance of being selected.
// If all values are zero or negative, the zero value of K is returned.
func (pm ProbabilityMap[K]) Pick() K {
	var total Probability
	for _, v := range pm {
		if v > 0 {
			total += v
		}
	}
	if total <= 0 {
		return zeroVal[K]()
	}

	r := r.Float64() * float64(total)
	var cumulative Probability
	for k, v := range pm {
		if v <= 0 {
			continue
		}
		cumulative += v
		if r < float64(cumulative) {
			return k
		}
	}
	return zeroVal[K]()
}
