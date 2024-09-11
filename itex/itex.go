package itex

import (
	"iter"
	"slices"
)

type PipeFn[V1, V2 any] func(V1) (V2, bool)

func Pipe[V1, V2 any](seq iter.Seq[V1], pipe PipeFn[V1, V2]) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		for {
			val, ok := next()
			if !ok {
				return
			}

			mapVal, ok := pipe(val)
			if !ok {
				continue
			}

			if !yield(mapVal) {
				return
			}
		}
	}
}

func Pipe2[K, V1, V2 any](seq iter.Seq2[K, V1], pipe PipeFn[V1, V2]) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()

		for {
			k, val, ok := next()
			if !ok {
				return
			}

			mapVal, ok := pipe(val)
			if !ok {
				continue
			}

			if !yield(k, mapVal) {
				return
			}
		}
	}
}

func FlatPipe[V1, V2 any](seq iter.Seq[[]V1], pipe PipeFn[V1, V2]) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		for {
			vals, ok := next()
			if !ok {
				return
			}

			innerSeq := slices.Values(vals)
			innerNext, innerStop := iter.Pull(innerSeq)
			defer innerStop()

			for {
				val, ok := innerNext()
				if !ok {
					return
				}

				mapVal, ok := pipe(val)
				if !ok {
					continue
				}

				if !yield(mapVal) {
					return
				}
			}
		}
	}
}

func Apply[V1, V2 any](slice []V1, pipe PipeFn[V1, V2]) []V2 {
	return slices.Collect(Pipe(slices.Values(slice), pipe))
}
