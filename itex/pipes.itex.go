package itex

func Filter[V any](pred func(V) bool) PipeFn[V, V] {
	return func(val V) (V, bool) {
		return val, pred(val)
	}
}

func Map[V1, V2 any](mapper func(V1) V2) PipeFn[V1, V2] {
	return func(val V1) (V2, bool) {
		return mapper(val), true
	}
}

func MapMaybe[V1, V2 any](mapper func(V1) (V2, bool)) PipeFn[V1, V2] {
	return func(val V1) (V2, bool) {
		return mapper(val)
	}
}
