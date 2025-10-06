package rng

import (
	"testing"
)

func TestRange_PanicOnNegativeMin(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative min")
		}
	}()
	Range(-1, 10)
}

func TestRange_PanicOnZeroMax(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero max")
		}
	}()
	Range(0, 0)
}

func TestRange_PanicOnNegativeMax(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative max")
		}
	}()
	Range(0, -1)
}

func TestRange_MinEqualsMax(t *testing.T) {
	result := Range(5, 5)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestRange_MinGreaterThanMax(t *testing.T) {
	result := Range(10, 5)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestRange_Int(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(1, 10)
		if result < 1 || result >= 10 {
			t.Errorf("Result %d out of range [1, 10)", result)
		}
	}
}

func TestRange_Uint(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(uint(1), uint(10))
		if result < 1 || result >= 10 {
			t.Errorf("Result %d out of range [1, 10)", result)
		}
	}
}

func TestRange_Float64(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(1.0, 10.0)
		if result < 1.0 || result >= 10.0 {
			t.Errorf("Result %f out of range [1.0, 10.0)", result)
		}
	}
}

func TestRange_Float32(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(float32(1.0), float32(10.0))
		if result < 1.0 || result >= 10.0 {
			t.Errorf("Result %f out of range [1.0, 10.0)", result)
		}
	}
}

func TestRange_Int8(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(int8(1), int8(10))
		if result < 1 || result >= 10 {
			t.Errorf("Result %d out of range [1, 10)", result)
		}
	}
}

func TestRange_Uint64(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Range(uint64(1), uint64(10))
		if result < 1 || result >= 10 {
			t.Errorf("Result %d out of range [1, 10)", result)
		}
	}
}
