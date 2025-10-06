package rng

import (
	"testing"
)

func TestProbabilityMap_Pick(t *testing.T) {
	t.Run("empty map returns zero value", func(t *testing.T) {
		pm := ProbabilityMap[string]{}
		result := pm.Pick()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("all zero or negative probabilities returns zero value", func(t *testing.T) {
		pm := ProbabilityMap[string]{
			"a": 0,
			"b": -0.5,
			"c": 0,
		}
		result := pm.Pick()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("single positive probability always returns that key", func(t *testing.T) {
		pm := ProbabilityMap[string]{
			"only": 1.0,
		}
		for range 10 {
			result := pm.Pick()
			if result != "only" {
				t.Errorf("expected 'only', got %q", result)
			}
		}
	})

	t.Run("picks from valid probabilities", func(t *testing.T) {
		pm := ProbabilityMap[string]{
			"a": 0.5,
			"b": 0.3,
			"c": 0.2,
		}

		// Run multiple times to check distribution roughly
		counts := make(map[string]int)
		iterations := 1000
		for range iterations {
			result := pm.Pick()
			counts[result]++
		}

		// Check that all keys were picked at least once
		for key := range pm {
			if counts[key] == 0 {
				t.Errorf("key %q was never picked", key)
			}
		}
	})

	t.Run("ignores zero and negative probabilities", func(t *testing.T) {
		pm := ProbabilityMap[int]{
			1: 0.5,
			2: 0,
			3: -0.1,
			4: 0.5,
		}

		// Run multiple times
		for range 50 {
			result := pm.Pick()
			if result != 1 && result != 4 {
				t.Errorf("expected 1 or 4, got %d", result)
			}
		}
	})
}

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
