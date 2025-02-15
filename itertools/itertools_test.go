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

func TestOnce(t *testing.T) {
	seq := Once(1)
	collected := slices.Collect(seq)
	if !slices.Equal(collected, []int{1}) {
		t.Fatalf("Once(1) != [1], got %v", collected)
	}
}

func TestSingle(t *testing.T) {
	s := []int{1, 2, 3}
	seq := slices.Values(s)
	{
		it := Single(seq)
		defer it.Stop()
		for range it.Seq() {
		}
		again := slices.Collect(it.Seq())
		if len(again) != 0 {
			t.Fatalf("Single([1, 2, 3]) not empty after use, again = %v", again)
		}
	}
	{
		it := Single(seq)
		defer it.Stop()
		for v := range it.Seq() {
			if v == 2 {
				break
			}
		}
		again := slices.Collect(it.Seq())
		if !slices.Equal(again, []int{3}) {
			t.Fatalf("Single([1, 2, 3]) not [3] after use with break, again = %v", again)
		}
	}
}

func TestDefer(t *testing.T) {
	r := 0
	d := 0
	for range Defer(func() {
		d++
	}) {
		if d != 0 {
			t.Errorf("Defer(f) f called in for loop, d = %v", d)
		}
		r++
	}
	if r != 1 {
		t.Errorf("Defer(f) loop body run not 1 time, r = %v", r)
	}
	if d != 1 {
		t.Errorf("Defer(f) f not called after for loop, d = %v", d)
	}
}
