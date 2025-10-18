package rng

import (
	"testing"
)

func TestNewLottery(t *testing.T) {
	// Test empty lottery
	lottery := NewLottery[string]()
	if lottery.Size() != 0 {
		t.Errorf("Expected empty lottery size 0, got %d", lottery.Size())
	}

	// Test lottery with initial items
	lottery = NewLottery("a", "b", "c")
	if lottery.Size() != 3 {
		t.Errorf("Expected lottery size 3, got %d", lottery.Size())
	}

	items := lottery.Items()
	for i, item := range items {
		expected := []string{"a", "b", "c"}[i]
		if item.Value != expected {
			t.Errorf("Expected item value %s, got %s", expected, item.Value)
		}
		if item.Weight != 1.0 {
			t.Errorf("Expected default weight 1.0, got %f", item.Weight)
		}
	}
}

func TestAppend(t *testing.T) {
	lottery := NewLottery[int]()
	lottery.Append(1, 2, 3)

	if lottery.Size() != 3 {
		t.Errorf("Expected size 3, got %d", lottery.Size())
	}

	items := lottery.Items()
	for i, item := range items {
		if item.Value != i+1 {
			t.Errorf("Expected value %d, got %d", i+1, item.Value)
		}
		if item.Weight != 1.0 {
			t.Errorf("Expected weight 1.0, got %f", item.Weight)
		}
	}
}

func TestAppendWithWeight(t *testing.T) {
	lottery := NewLottery[string]()
	lottery.AppendWeight(2.5, "heavy")
	lottery.AppendWeight(0.5, "light")

	items := lottery.Items()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Weight != 2.5 {
		t.Errorf("Expected weight 2.5, got %f", items[0].Weight)
	}
	if items[1].Weight != 0.5 {
		t.Errorf("Expected weight 0.5, got %f", items[1].Weight)
	}
}

func TestDraw(t *testing.T) {
	// Test empty lottery
	lottery := NewLottery[string]()
	result := lottery.Draw()
	if result != "" {
		t.Errorf("Expected zero value for empty lottery, got %s", result)
	}

	// Test single item
	lottery.Append("only")
	result = lottery.Draw()
	if result != "only" {
		t.Errorf("Expected 'only', got %s", result)
	}

	// Test multiple items with weights
	lottery = NewLottery[string]()
	lottery.AppendWeight(1.0, "a")
	lottery.AppendWeight(2.0, "b")
	lottery.AppendWeight(3.0, "c")

	// Draw multiple times to check distribution roughly
	counts := make(map[string]int)
	for i := 0; i < 1000; i++ {
		result := lottery.Draw()
		counts[result]++
	}

	// c should be drawn most often, a least often
	if counts["c"] <= counts["b"] || counts["b"] <= counts["a"] {
		t.Logf("Distribution may not match weights exactly: a=%d, b=%d, c=%d", counts["a"], counts["b"], counts["c"])
	}
}

func TestDrawWithZeroWeights(t *testing.T) {
	lottery := NewLottery[string]()
	lottery.AppendWeight(0, "zero1")
	lottery.AppendWeight(0, "zero2")

	// Should still return one of the items as fallback
	result := lottery.Draw()
	if result != "zero1" && result != "zero2" {
		t.Errorf("Expected fallback to return one of the zero-weight items, got %s", result)
	}
}

func TestDrawCount(t *testing.T) {
	lottery := NewLottery("test")

	// Initial draw count should be 0
	items := lottery.Items()
	if items[0].DrawCount != 0 {
		t.Errorf("Expected initial draw count 0, got %d", items[0].DrawCount)
	}

	// After drawing, count should increment
	lottery.Draw()
	items = lottery.Items()
	if items[0].DrawCount != 1 {
		t.Errorf("Expected draw count 1 after one draw, got %d", items[0].DrawCount)
	}
}

func TestClear(t *testing.T) {
	lottery := NewLottery("a", "b", "c")
	if lottery.Size() != 3 {
		t.Errorf("Expected initial size 3, got %d", lottery.Size())
	}

	lottery.Clear()
	if lottery.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", lottery.Size())
	}

	items := lottery.Items()
	if len(items) != 0 {
		t.Errorf("Expected empty items after clear, got %d items", len(items))
	}
}

func TestSize(t *testing.T) {
	lottery := NewLottery[int]()
	if lottery.Size() != 0 {
		t.Errorf("Expected initial size 0, got %d", lottery.Size())
	}

	lottery.Append(1)
	if lottery.Size() != 1 {
		t.Errorf("Expected size 1 after append, got %d", lottery.Size())
	}

	lottery.Append(2, 3)
	if lottery.Size() != 3 {
		t.Errorf("Expected size 3 after appending 2 more, got %d", lottery.Size())
	}
}

func TestItems(t *testing.T) {
	lottery := NewLottery[string]()
	lottery.AppendWeight(1.5, "first")
	lottery.AppendWeight(2.5, "second")

	items := lottery.Items()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Value != "first" || items[0].Weight != 1.5 {
		t.Errorf("Expected first item: value='first', weight=1.5, got value='%s', weight=%f", items[0].Value, items[0].Weight)
	}

	if items[1].Value != "second" || items[1].Weight != 2.5 {
		t.Errorf("Expected second item: value='second', weight=2.5, got value='%s', weight=%f", items[1].Value, items[1].Weight)
	}

	// Test that modifying returned items doesn't affect original
	items[0].Weight = 999
	originalItems := lottery.Items()
	if originalItems[0].Weight == 999 {
		t.Error("Items() should return a copy, not reference to original data")
	}
}

func TestConcurrency(t *testing.T) {
	lottery := NewLottery[int]()

	// Test concurrent appends
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(val int) {
			lottery.Append(val)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if lottery.Size() != 10 {
		t.Errorf("Expected size 10 after concurrent appends, got %d", lottery.Size())
	}

	// Test concurrent draws
	for i := 0; i < 10; i++ {
		go func() {
			lottery.Draw()
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestAppendWeights(t *testing.T) {
	// Test with empty map
	lottery := NewLottery[string]()
	lottery.AppendWeights(map[float64][]string{})
	if lottery.Size() != 0 {
		t.Errorf("Expected size 0 after appending empty map, got %d", lottery.Size())
	}

	// Test with single weight and single value
	lottery = NewLottery[string]()
	lottery.AppendWeights(map[float64][]string{
		2.5: {"item1"},
	})
	if lottery.Size() != 1 {
		t.Errorf("Expected size 1, got %d", lottery.Size())
	}
	items := lottery.Items()
	if items[0].Value != "item1" || items[0].Weight != 2.5 {
		t.Errorf("Expected item1 with weight 2.5, got %s with weight %f", items[0].Value, items[0].Weight)
	}

	// Test with single weight and multiple values
	lottery = NewLottery[string]()
	lottery.AppendWeights(map[float64][]string{
		1.5: {"a", "b", "c"},
	})
	if lottery.Size() != 3 {
		t.Errorf("Expected size 3, got %d", lottery.Size())
	}
	items = lottery.Items()
	expectedValues := []string{"a", "b", "c"}
	foundValues := make(map[string]bool)
	for _, item := range items {
		if item.Weight != 1.5 {
			t.Errorf("Expected weight 1.5 for all items, got %f for %s", item.Weight, item.Value)
		}
		foundValues[item.Value] = true
	}
	for _, expected := range expectedValues {
		if !foundValues[expected] {
			t.Errorf("Expected to find value %s but didn't", expected)
		}
	}

	// Test with multiple weights and multiple values
	lottery2 := NewLottery[int]()
	lottery2.AppendWeights(map[float64][]int{
		1.0: {1, 2},
		2.0: {3, 4, 5},
		3.0: {6},
	})
	if lottery2.Size() != 6 {
		t.Errorf("Expected size 6, got %d", lottery2.Size())
	}
	items2 := lottery2.Items()
	weightMap := make(map[int]float64)
	for _, item := range items2 {
		weightMap[item.Value] = item.Weight
	}
	if weightMap[1] != 1.0 || weightMap[2] != 1.0 {
		t.Errorf("Expected weight 1.0 for values 1 and 2")
	}
	if weightMap[3] != 2.0 || weightMap[4] != 2.0 || weightMap[5] != 2.0 {
		t.Errorf("Expected weight 2.0 for values 3, 4, and 5")
	}
	if weightMap[6] != 3.0 {
		t.Errorf("Expected weight 3.0 for value 6")
	}

	// Test with zero and negative weights
	lottery = NewLottery[string]()
	lottery.AppendWeights(map[float64][]string{
		0.0:  {"zero"},
		-1.0: {"negative"},
		1.0:  {"positive"},
	})
	if lottery.Size() != 3 {
		t.Errorf("Expected size 3, got %d", lottery.Size())
	}
	items = lottery.Items()
	for _, item := range items {
		switch item.Value {
		case "zero":
			if item.Weight != 0.0 {
				t.Errorf("Expected weight 0.0 for 'zero', got %f", item.Weight)
			}
		case "negative":
			if item.Weight != -1.0 {
				t.Errorf("Expected weight -1.0 for 'negative', got %f", item.Weight)
			}
		case "positive":
			if item.Weight != 1.0 {
				t.Errorf("Expected weight 1.0 for 'positive', got %f", item.Weight)
			}
		}
	}
}

func TestAppendWeightsChaining(t *testing.T) {
	lottery := NewLottery[string]()
	result := lottery.AppendWeights(map[float64][]string{
		1.0: {"a"},
	}).AppendWeights(map[float64][]string{
		2.0: {"b"},
	})

	if result != lottery {
		t.Error("AppendWeights should return the same lottery instance for chaining")
	}

	if lottery.Size() != 2 {
		t.Errorf("Expected size 2 after chaining, got %d", lottery.Size())
	}
}

func TestAppendWeightsWithExistingItems(t *testing.T) {
	lottery := NewLottery("existing")
	lottery.AppendWeights(map[float64][]string{
		2.0: {"new1", "new2"},
	})

	if lottery.Size() != 3 {
		t.Errorf("Expected size 3 (1 existing + 2 new), got %d", lottery.Size())
	}

	items := lottery.Items()
	found := make(map[string]float64)
	for _, item := range items {
		found[item.Value] = item.Weight
	}

	if found["existing"] != 1.0 {
		t.Errorf("Expected existing item to have weight 1.0, got %f", found["existing"])
	}
	if found["new1"] != 2.0 || found["new2"] != 2.0 {
		t.Errorf("Expected new items to have weight 2.0")
	}
}

func TestAppendWeightsConcurrency(t *testing.T) {
	lottery := NewLottery[int]()
	done := make(chan bool, 5)

	// Concurrent AppendWeights calls
	for i := 0; i < 5; i++ {
		go func(val int) {
			lottery.AppendWeights(map[float64][]int{
				float64(val): {val, val * 10},
			})
			done <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	if lottery.Size() != 10 {
		t.Errorf("Expected size 10 after concurrent appends, got %d", lottery.Size())
	}
}

func TestAppendWeightsEmptySlices(t *testing.T) {
	lottery := NewLottery[string]()
	lottery.AppendWeights(map[float64][]string{
		1.0: {},
		2.0: {"valid"},
		3.0: {},
	})

	if lottery.Size() != 1 {
		t.Errorf("Expected size 1 (only non-empty slice), got %d", lottery.Size())
	}

	items := lottery.Items()
	if items[0].Value != "valid" || items[0].Weight != 2.0 {
		t.Errorf("Expected 'valid' with weight 2.0, got %s with weight %f", items[0].Value, items[0].Weight)
	}
}
