package rng

var P50 = Probability(0.5)

// Probability represents a probability value between 0 and 1 (inclusive).
type Probability float64

// Check returns true with the probability specified by p.
//
// If p is less than or equal to 0, it always returns false.
// If p is greater than or equal to 1, it always returns true.
// For values between 0 and 1, it returns true with probability p using a random number generator.
func (p Probability) Check() bool {
	if p <= 0 {
		return false
	}
	if p >= 1 {
		return true
	}
	return r.Float64() < float64(p)
}
