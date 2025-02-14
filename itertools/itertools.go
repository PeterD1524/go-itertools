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
