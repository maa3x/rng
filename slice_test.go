package rng

import (
	"fmt"
	"slices"
	"testing"
)

func TestPick(t *testing.T) {
	t.Run("empty slice returns zero value", func(t *testing.T) {
		var empty []int
		result := Pick(empty)
		if result != 0 {
			t.Errorf("expected 0 for empty int slice, got %d", result)
		}

		var emptyStr []string
		resultStr := Pick(emptyStr)
		if resultStr != "" {
			t.Errorf("expected empty string for empty string slice, got %q", resultStr)
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		slice := []int{42}
		result := Pick(slice)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("multiple elements slice contains result", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := Pick(slice)

		found := slices.Contains(slice, result)
		if !found {
			t.Errorf("result %d not found in slice %v", result, slice)
		}
	})

	t.Run("string slice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		result := Pick(slice)

		found := slices.Contains(slice, result)
		if !found {
			t.Errorf("result %q not found in slice %v", result, slice)
		}
	})
}

func TestPickN(t *testing.T) {
	t.Run("returns nil for n <= 0", func(t *testing.T) {
		slice := []int{1, 2, 3}

		result := PickN(slice, 0)
		if result != nil {
			t.Errorf("expected nil for n=0, got %v", result)
		}

		result = PickN(slice, -1)
		if result != nil {
			t.Errorf("expected nil for n=-1, got %v", result)
		}
	})

	t.Run("returns nil for empty slice", func(t *testing.T) {
		var empty []int
		result := PickN(empty, 5)
		if result != nil {
			t.Errorf("expected nil for empty slice, got %v", result)
		}
	})

	t.Run("returns correct length", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		result := PickN(slice, 3)
		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		result = PickN(slice, 1)
		if len(result) != 1 {
			t.Errorf("expected length 1, got %d", len(result))
		}
	})

	t.Run("all elements are from original slice", func(t *testing.T) {
		slice := []int{10, 20, 30}
		result := PickN(slice, 5)

		for _, v := range result {
			found := slices.Contains(slice, v)
			if !found {
				t.Errorf("element %d not found in original slice %v", v, slice)
			}
		}
	})

	t.Run("can pick more elements than slice length", func(t *testing.T) {
		slice := []string{"a", "b"}
		result := PickN(slice, 5)

		if len(result) != 5 {
			t.Errorf("expected length 5, got %d", len(result))
		}

		// All elements should still be from the original slice
		for _, v := range result {
			if v != "a" && v != "b" {
				t.Errorf("unexpected element %q, expected 'a' or 'b'", v)
			}
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		slice := []int{42}
		result := PickN(slice, 3)

		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		for _, v := range result {
			if v != 42 {
				t.Errorf("expected all elements to be 42, got %d", v)
			}
		}
	})

	t.Run("works with different types", func(t *testing.T) {
		floatSlice := []float64{1.1, 2.2, 3.3}
		result := PickN(floatSlice, 2)

		if len(result) != 2 {
			t.Errorf("expected length 2, got %d", len(result))
		}

		for _, v := range result {
			found := slices.Contains(floatSlice, v)
			if !found {
				t.Errorf("element %f not found in original slice %v", v, floatSlice)
			}
		}
	})
}

func TestPickNDistinct(t *testing.T) {
	t.Run("returns nil for n <= 0", func(t *testing.T) {
		slice := []int{1, 2, 3}

		result := PickNDistinct(slice, 0)
		if result != nil {
			t.Errorf("expected nil for n=0, got %v", result)
		}

		result = PickNDistinct(slice, -1)
		if result != nil {
			t.Errorf("expected nil for n=-1, got %v", result)
		}
	})

	t.Run("panics when n > slice length", func(t *testing.T) {
		slice := []int{1, 2, 3}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic when n > slice length")
			}
		}()

		PickNDistinct(slice, 5)
	})

	t.Run("returns correct length", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		result := PickNDistinct(slice, 3)
		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		result = PickNDistinct(slice, 1)
		if len(result) != 1 {
			t.Errorf("expected length 1, got %d", len(result))
		}
	})

	t.Run("all elements are distinct", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := PickNDistinct(slice, 3)

		// Check for duplicates
		seen := make(map[int]bool)
		for _, v := range result {
			if seen[v] {
				t.Errorf("found duplicate element %d in result %v", v, result)
			}
			seen[v] = true
		}
	})

	t.Run("all elements are from original slice", func(t *testing.T) {
		slice := []int{10, 20, 30, 40}
		result := PickNDistinct(slice, 3)

		for _, v := range result {
			found := slices.Contains(slice, v)
			if !found {
				t.Errorf("element %d not found in original slice %v", v, slice)
			}
		}
	})

	t.Run("can pick all elements", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		result := PickNDistinct(slice, 3)

		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		// All original elements should be present
		originalMap := make(map[string]bool)
		for _, v := range slice {
			originalMap[v] = true
		}

		resultMap := make(map[string]bool)
		for _, v := range result {
			resultMap[v] = true
		}

		for k := range originalMap {
			if !resultMap[k] {
				t.Errorf("original element %q not found in result", k)
			}
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		slice := []int{42}
		result := PickNDistinct(slice, 1)

		if len(result) != 1 {
			t.Errorf("expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("expected element to be 42, got %d", result[0])
		}
	})

	t.Run("works with different types", func(t *testing.T) {
		floatSlice := []float64{1.1, 2.2, 3.3, 4.4}
		result := PickNDistinct(floatSlice, 2)

		if len(result) != 2 {
			t.Errorf("expected length 2, got %d", len(result))
		}

		// Check distinctness
		if result[0] == result[1] {
			t.Errorf("found duplicate elements in result %v", result)
		}

		// Check elements are from original
		for _, v := range result {
			found := slices.Contains(floatSlice, v)
			if !found {
				t.Errorf("element %f not found in original slice %v", v, floatSlice)
			}
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic for empty slice with n > 0")
			}
		}()

		PickNDistinct(empty, 1)
	})
}

func TestPickNUnique(t *testing.T) {
	t.Run("returns error for n <= 0", func(t *testing.T) {
		slice := []int{1, 2, 3}

		_, err := PickNUnique(slice, 0)
		if err == nil {
			t.Error("expected error for n=0")
		}

		_, err = PickNUnique(slice, -1)
		if err == nil {
			t.Error("expected error for n=-1")
		}
	})

	t.Run("returns error when n > unique elements", func(t *testing.T) {
		slice := []int{1, 1, 2, 2, 3} // 3 unique elements

		_, err := PickNUnique(slice, 4)
		if err == nil {
			t.Error("expected error when n > unique elements")
		}
	})

	t.Run("returns correct length", func(t *testing.T) {
		slice := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5} // 5 unique elements

		result, err := PickNUnique(slice, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}
	})

	t.Run("all elements are distinct", func(t *testing.T) {
		slice := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

		result, err := PickNUnique(slice, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Check for duplicates
		seen := make(map[int]bool)
		for _, v := range result {
			if seen[v] {
				t.Errorf("found duplicate element %d in result %v", v, result)
			}
			seen[v] = true
		}
	})

	t.Run("all elements are from original unique set", func(t *testing.T) {
		slice := []int{1, 1, 2, 2, 3, 3, 4, 4}
		unique := []int{1, 2, 3, 4}

		result, err := PickNUnique(slice, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		for _, v := range result {
			found := slices.Contains(unique, v)
			if !found {
				t.Errorf("element %d not found in unique set %v", v, unique)
			}
		}
	})

	t.Run("can pick all unique elements", func(t *testing.T) {
		slice := []string{"a", "a", "b", "b", "c", "c"}

		result, err := PickNUnique(slice, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		// Check that all unique elements are present
		expectedUnique := map[string]bool{"a": true, "b": true, "c": true}
		resultMap := make(map[string]bool)
		for _, v := range result {
			resultMap[v] = true
		}

		for k := range expectedUnique {
			if !resultMap[k] {
				t.Errorf("unique element %q not found in result", k)
			}
		}
	})

	t.Run("slice with no duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		result, err := PickNUnique(slice, 3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 3 {
			t.Errorf("expected length 3, got %d", len(result))
		}

		// Check distinctness
		seen := make(map[int]bool)
		for _, v := range result {
			if seen[v] {
				t.Errorf("found duplicate element %d in result %v", v, result)
			}
			seen[v] = true
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		slice := []int{42, 42, 42}

		result, err := PickNUnique(slice, 1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 1 {
			t.Errorf("expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("expected element to be 42, got %d", result[0])
		}
	})

	t.Run("works with different types", func(t *testing.T) {
		floatSlice := []float64{1.1, 1.1, 2.2, 2.2, 3.3, 3.3, 4.4}

		result, err := PickNUnique(floatSlice, 2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 2 {
			t.Errorf("expected length 2, got %d", len(result))
		}

		// Check distinctness
		if result[0] == result[1] {
			t.Errorf("found duplicate elements in result %v", result)
		}

		// Check elements are from unique set
		unique := []float64{1.1, 2.2, 3.3, 4.4}
		for _, v := range result {
			found := slices.Contains(unique, v)
			if !found {
				t.Errorf("element %f not found in unique set %v", v, unique)
			}
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int

		_, err := PickNUnique(empty, 1)
		if err == nil {
			t.Error("expected error for empty slice with n > 0")
		}
	})

	t.Run("slice with all duplicate elements", func(t *testing.T) {
		slice := []string{"same", "same", "same", "same"}

		result, err := PickNUnique(slice, 1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(result) != 1 {
			t.Errorf("expected length 1, got %d", len(result))
		}

		if result[0] != "same" {
			t.Errorf("expected element to be 'same', got %q", result[0])
		}

		// Should error when asking for more than 1 unique element
		_, err = PickNUnique(slice, 2)
		if err == nil {
			t.Error("expected error when n > unique elements")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var empty []int
		originalLen := len(empty)
		Shuffle(empty)
		if len(empty) != originalLen {
			t.Errorf("expected length to remain %d, got %d", originalLen, len(empty))
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		slice := []int{42}
		original := make([]int, len(slice))
		copy(original, slice)

		Shuffle(slice)

		if len(slice) != len(original) {
			t.Errorf("expected length to remain %d, got %d", len(original), len(slice))
		}
		if slice[0] != original[0] {
			t.Errorf("single element should remain unchanged, expected %d, got %d", original[0], slice[0])
		}
	})

	t.Run("preserves all elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		original := make([]int, len(slice))
		copy(original, slice)

		Shuffle(slice)

		// Check length is preserved
		if len(slice) != len(original) {
			t.Errorf("expected length to remain %d, got %d", len(original), len(slice))
		}

		// Check all elements are still present
		originalMap := make(map[int]int)
		for _, v := range original {
			originalMap[v]++
		}

		shuffledMap := make(map[int]int)
		for _, v := range slice {
			shuffledMap[v]++
		}

		for k, v := range originalMap {
			if shuffledMap[k] != v {
				t.Errorf("element %d count mismatch: expected %d, got %d", k, v, shuffledMap[k])
			}
		}
	})

	t.Run("modifies slice in place", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		slicePtr := &slice[0]

		Shuffle(slice)

		// Verify we're still working with the same underlying array
		if &slice[0] != slicePtr {
			t.Error("slice should be modified in place, not reallocated")
		}
	})

	t.Run("produces different arrangements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
		original := make([]int, len(slice))
		copy(original, slice)

		// Run shuffle multiple times and check if we get at least one different arrangement
		arrangements := make(map[string]bool)
		maxAttempts := 100

		for range maxAttempts {
			testSlice := make([]int, len(original))
			copy(testSlice, original)
			Shuffle(testSlice)

			// Convert to string for comparison
			arrangement := fmt.Sprintf("%v", testSlice)
			arrangements[arrangement] = true
		}

		// With 8 elements, we should get multiple different arrangements
		if len(arrangements) < 2 {
			t.Errorf("expected multiple different arrangements, got only %d unique arrangements", len(arrangements))
		}
	})

	t.Run("works with different types", func(t *testing.T) {
		// Test with strings
		stringSlice := []string{"apple", "banana", "cherry", "date"}
		originalStrings := make([]string, len(stringSlice))
		copy(originalStrings, stringSlice)

		Shuffle(stringSlice)

		if len(stringSlice) != len(originalStrings) {
			t.Errorf("string slice length changed: expected %d, got %d", len(originalStrings), len(stringSlice))
		}

		// Check all strings are preserved
		for _, orig := range originalStrings {
			found := slices.Contains(stringSlice, orig)
			if !found {
				t.Errorf("string %q not found in shuffled slice", orig)
			}
		}

		// Test with floats
		floatSlice := []float64{1.1, 2.2, 3.3, 4.4}
		originalFloats := make([]float64, len(floatSlice))
		copy(originalFloats, floatSlice)

		Shuffle(floatSlice)

		if len(floatSlice) != len(originalFloats) {
			t.Errorf("float slice length changed: expected %d, got %d", len(originalFloats), len(floatSlice))
		}

		// Check all floats are preserved
		for _, orig := range originalFloats {
			found := slices.Contains(floatSlice, orig)
			if !found {
				t.Errorf("float %f not found in shuffled slice", orig)
			}
		}
	})

	t.Run("two element slice", func(t *testing.T) {
		// Run multiple times to see both possible arrangements
		arrangements := make(map[string]bool)
		for range 20 {
			testSlice := []int{1, 2}
			Shuffle(testSlice)
			arrangement := fmt.Sprintf("%v", testSlice)
			arrangements[arrangement] = true
		}

		// Should have at most 2 arrangements: [1 2] and [2 1]
		if len(arrangements) > 2 {
			t.Errorf("expected at most 2 arrangements for 2-element slice, got %d", len(arrangements))
		}

		// Verify both elements are always present
		for arrangement := range arrangements {
			if arrangement != "[1 2]" && arrangement != "[2 1]" {
				t.Errorf("unexpected arrangement: %s", arrangement)
			}
		}
	})

	t.Run("slice with duplicate elements", func(t *testing.T) {
		slice := []int{1, 1, 2, 2, 3, 3}
		original := make([]int, len(slice))
		copy(original, slice)

		Shuffle(slice)

		// Check length is preserved
		if len(slice) != len(original) {
			t.Errorf("expected length to remain %d, got %d", len(original), len(slice))
		}

		// Count occurrences of each element
		originalCount := make(map[int]int)
		for _, v := range original {
			originalCount[v]++
		}

		shuffledCount := make(map[int]int)
		for _, v := range slice {
			shuffledCount[v]++
		}

		// Verify counts match
		for k, v := range originalCount {
			if shuffledCount[k] != v {
				t.Errorf("element %d count mismatch: expected %d, got %d", k, v, shuffledCount[k])
			}
		}
	})
}
