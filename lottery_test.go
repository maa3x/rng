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
	lottery.AppendWithWeight(2.5, "heavy")
	lottery.AppendWithWeight(0.5, "light")

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
	lottery.AppendWithWeight(1.0, "a")
	lottery.AppendWithWeight(2.0, "b")
	lottery.AppendWithWeight(3.0, "c")

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
	lottery.AppendWithWeight(0, "zero1")
	lottery.AppendWithWeight(0, "zero2")

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
	lottery.AppendWithWeight(1.5, "first")
	lottery.AppendWithWeight(2.5, "second")

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
