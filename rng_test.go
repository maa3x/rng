package rng

import (
	"testing"
)

func TestN(t *testing.T) {
	// Test with deterministic seed for reproducible results
	tests := []struct {
		name string
		n    int
	}{
		{"n=1", 1},
		{"n=10", 10},
		{"n=100", 100},
		{"n=1000", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result := N(tt.n)
				if result < 0 || result >= tt.n {
					t.Errorf("N(%d) = %d, want value in range [0, %d)", tt.n, result, tt.n)
				}
			}
		})
	}
}

func TestN_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for n <= 0")
		}
	}()
	N(0)
}

func TestN_NegativePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative n")
		}
	}()
	N(-1)
}

func TestN_DifferentIntTypes(t *testing.T) {
	// Test with different integer types
	result8 := N[int8](10)
	if result8 < 0 || result8 >= 10 {
		t.Errorf("N[int8](10) = %d, want value in range [0, 10)", result8)
	}

	result16 := N[int16](100)
	if result16 < 0 || result16 >= 100 {
		t.Errorf("N[int16](100) = %d, want value in range [0, 100)", result16)
	}

	result32 := N[int32](1000)
	if result32 < 0 || result32 >= 1000 {
		t.Errorf("N[int32](1000) = %d, want value in range [0, 1000)", result32)
	}

	result64 := N[int64](10000)
	if result64 < 0 || result64 >= 10000 {
		t.Errorf("N[int64](10000) = %d, want value in range [0, 10000)", result64)
	}
}

func TestNum(t *testing.T) {
	// Test float types
	t.Run("float32", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result := Num[float32]()
			if result < 0.0 || result >= 1.0 {
				t.Errorf("Num[float32]() = %f, want value in range [0.0, 1.0)", result)
			}
		}
	})

	t.Run("float64", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result := Num[float64]()
			if result < 0.0 || result >= 1.0 {
				t.Errorf("Num[float64]() = %f, want value in range [0.0, 1.0)", result)
			}
		}
	})

	// Test integer types
	t.Run("int", func(t *testing.T) {
		result := Num[int]()
		// Just verify it returns without error
		_ = result
	})

	t.Run("int8", func(t *testing.T) {
		result := Num[int8]()
		_ = result
	})

	t.Run("int16", func(t *testing.T) {
		result := Num[int16]()
		_ = result
	})

	t.Run("int32", func(t *testing.T) {
		result := Num[int32]()
		_ = result
	})

	t.Run("int64", func(t *testing.T) {
		result := Num[int64]()
		_ = result
	})

	t.Run("uint", func(t *testing.T) {
		result := Num[uint]()
		_ = result
	})

	t.Run("uint8", func(t *testing.T) {
		result := Num[uint8]()
		_ = result
	})

	t.Run("uint16", func(t *testing.T) {
		result := Num[uint16]()
		_ = result
	})

	t.Run("uint32", func(t *testing.T) {
		result := Num[uint32]()
		_ = result
	})

	t.Run("uint64", func(t *testing.T) {
		result := Num[uint64]()
		_ = result
	})
}

func TestNum_Randomness(t *testing.T) {
	// Test that consecutive calls produce different values (probabilistically)
	const iterations = 1000

	t.Run("float64_randomness", func(t *testing.T) {
		values := make(map[float64]bool)
		duplicates := 0

		for i := 0; i < iterations; i++ {
			val := Num[float64]()
			if values[val] {
				duplicates++
			}
			values[val] = true
		}

		// Allow some duplicates but expect mostly unique values
		if duplicates > iterations/10 {
			t.Errorf("Too many duplicate values: %d out of %d", duplicates, iterations)
		}
	})

	t.Run("uint64_randomness", func(t *testing.T) {
		values := make(map[uint64]bool)
		duplicates := 0

		for i := 0; i < iterations; i++ {
			val := Num[uint64]()
			if values[val] {
				duplicates++
			}
			values[val] = true
		}

		// Allow some duplicates but expect mostly unique values
		if duplicates > iterations/10 {
			t.Errorf("Too many duplicate values: %d out of %d", duplicates, iterations)
		}
	})
}
