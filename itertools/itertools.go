package itertools

import "iter"

func Map[V any, B any](seq iter.Seq[V], f func(V) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}
