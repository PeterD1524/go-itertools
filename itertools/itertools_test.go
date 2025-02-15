package itertools

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3}
	f := func(x int) int { return 2 * x }
	mapped := slices.Collect(Map(slices.Values(s), f))
	want := []int{2, 4, 6}
	if !slices.Equal(mapped, want) {
		t.Fatalf("seq = %v, f(x) = 2 * x, Map(seq, f) = %v, want = %v", s, mapped, want)
	}
}

func TestAllEmpty(t *testing.T) {
	seq := []struct{}{}
	f := func(x struct{}) bool { return false }
	if !All(slices.Values(seq), f) {
		t.Fatalf("All([]) = false")
	}
}

func TestAllSimple(t *testing.T) {
	s := []int{1, 2, 3}
	if f := func(x int) bool { return x > 0 }; !All(slices.Values(s), f) {
		t.Errorf("f(x) = x > 0, All([1, 2, 3], f) = false")
	}
	if f := func(x int) bool { return x > 2 }; All(slices.Values(s), f) {
		t.Fatalf("f(x) = x > 2, All([1, 2, 3], f) = true")
	}
}

func TestAllShortCircuit(t *testing.T) {
	s := []int{1, 2, 3}
	count := 0
	f := func(x int) bool {
		count++
		return x != 2
	}
	if All(slices.Values(s), f) {
		t.Fatalf("f(x) = x != 2, All([1, 2, 3], f) = true")
	}
	if count != 2 {
		t.Fatalf("f(x) = x != 2, All([1, 2, 3], f) called f %v times, not 2", count)
	}
}
