package itertools

import "iter"

func All[V any](seq iter.Seq[V], f func(V) bool) bool {
	for b := range Map(seq, f) {
		if !b {
			return false
		}
	}
	return true
}

func Map[V any, B any](seq iter.Seq[V], f func(V) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

type Iterator[V any] struct {
	next func() (V, bool)
	stop func()
	seq  iter.Seq[V]
}

func (it *Iterator[V]) Next() (V, bool) {
	return it.next()
}

func (it *Iterator[V]) Stop() {
	it.stop()
}

func (it *Iterator[V]) Seq() iter.Seq[V] {
	return it.seq
}

// Single converts a possible multi-use iterator into a single-use resumable iterator.
// Must `defer it.Stop()` or fully consume the iterator, otherwise you have a goroutine (coroutine created by iter.Pull) leak.
// See https://pkg.go.dev/iter#hdr-Single_Use_Iterators.
// Goroutine leak detector: https://github.com/uber-go/goleak
func Single[V any](seq iter.Seq[V]) Iterator[V] {
	next, stop := iter.Pull(seq)
	return Iterator[V]{next, stop, func(yield func(V) bool) {
		for {
			v, ok := next()
			if !ok || !yield(v) {
				break
			}
		}
	}}
}
