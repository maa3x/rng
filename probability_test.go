package rng

import (
	"testing"
)

func TestProbability_Check(t *testing.T) {
	t.Run("zero probability always returns false", func(t *testing.T) {
		p := Probability(0)
		for range 10 {
			if p.Check() {
				t.Error("expected false for zero probability")
			}
		}
	})

	t.Run("negative probability always returns false", func(t *testing.T) {
		p := Probability(-0.5)
		for range 10 {
			if p.Check() {
				t.Error("expected false for negative probability")
			}
		}
	})

	t.Run("probability of 1 always returns true", func(t *testing.T) {
		p := Probability(1.0)
		for range 10 {
			if !p.Check() {
				t.Error("expected true for probability of 1")
			}
		}
	})

	t.Run("probability greater than 1 always returns true", func(t *testing.T) {
		p := Probability(1.5)
		for range 10 {
			if !p.Check() {
				t.Error("expected true for probability greater than 1")
			}
		}
	})

	t.Run("probability between 0 and 1 returns both true and false", func(t *testing.T) {
		p := Probability(0.5)
		trueCount := 0
		falseCount := 0
		iterations := 1000

		for range iterations {
			if p.Check() {
				trueCount++
			} else {
				falseCount++
			}
		}

		if trueCount == 0 {
			t.Error("expected at least some true results for probability 0.5")
		}
		if falseCount == 0 {
			t.Error("expected at least some false results for probability 0.5")
		}
	})

	t.Run("very small positive probability sometimes returns true", func(t *testing.T) {
		p := Probability(0.01)
		hasTrue := false
		hasFalse := false

		for range 1000 {
			if p.Check() {
				hasTrue = true
			} else {
				hasFalse = true
			}
			if hasTrue && hasFalse {
				break
			}
		}

		if !hasFalse {
			t.Error("expected at least some false results for small probability")
		}
	})
}
