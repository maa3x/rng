package rng

import (
	"testing"
)

func TestWeightMap_Pick(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		wm := WeightMap[string]{}
		result := wm.Pick()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("all zero weights", func(t *testing.T) {
		wm := WeightMap[string]{
			"a": 0,
			"b": 0,
			"c": 0,
		}
		result := wm.Pick()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("single item", func(t *testing.T) {
		wm := WeightMap[string]{
			"only": 10,
		}
		result := wm.Pick()
		if result != "only" {
			t.Errorf("expected 'only', got %q", result)
		}
	})

	t.Run("multiple items with equal weights", func(t *testing.T) {
		wm := WeightMap[int]{
			1: 5,
			2: 5,
			3: 5,
		}
		t.Logf("WeightMap: %+v", wm)
		// Run multiple times to check distribution
		results := make(map[int]int)
		for range 10000 {
			result := wm.Pick()
			results[result]++
		}
		t.Logf("selection distribution: %+v", results)
		// All keys should be selected at least once
		for key := 1; key <= 3; key++ {
			if results[key] == 0 {
				t.Errorf("key %d was never selected", key)
			}
		}
	})

	t.Run("different weights distribution", func(t *testing.T) {
		wm := WeightMap[string]{
			"high":   80,
			"medium": 20,
			"low":    10,
		}
		t.Logf("WeightMap: %+v", wm)
		// Run multiple times to check distribution
		results := make(map[string]int)
		for range 1000000 {
			result := wm.Pick()
			results[result]++
		}
		t.Logf("selection distribution: %+v", results)
		// High weight should be selected most often
		if results["high"] <= results["medium"] || results["high"] <= results["low"] {
			t.Error("high weight item should be selected most frequently")
		}
		if results["medium"] <= results["low"] {
			t.Error("medium weight item should be selected more than low weight item")
		}
	})
}
